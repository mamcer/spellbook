# Navidrome

Self-hosted music server with a web UI and the Subsonic API (so any
Subsonic client works on phones).

[https://www.navidrome.org/docs/](https://www.navidrome.org/docs/)

[https://www.navidrome.org/apps/](https://www.navidrome.org/apps/) (client directory, with last-update dates)

## try it before installing

public demo: [https://demo.navidrome.org](https://demo.navidrome.org) (user `demo` / password `demo`)

throwaway instance against your own library, read-only, data in `/tmp`:

```bash
ND_VER=$(curl -s https://api.github.com/repos/navidrome/navidrome/releases/latest \
  | jq -r .tag_name | sed 's/^v//')

mkdir -p /tmp/navidrome-trial && chmod 777 /tmp/navidrome-trial

docker run --rm --name navidrome-trial \
  -p 4533:4533 \
  --user 65534:65534 \
  -v /path/to/music:/music:ro \
  -v /tmp/navidrome-trial:/data \
  deluan/navidrome:$ND_VER
# http://<server-ip>:4533 -> the first user to log in becomes admin
```

cleanup (files in /tmp belong to `nobody`, so delete them from a container):

```bash
docker stop navidrome-trial
docker run --rm -v /tmp:/t alpine:3.20 rm -rf /t/navidrome-trial
```

## keep the music untouched

Navidrome never edits tags or writes to the music folder; its DB and
cache go to the data folder. Belt and braces when the music is a backup:

- mount it read-only (`:ro` / `readOnly: true`), enforced by the kernel
- run as `nobody` (65534), not as the owner of the files
- files not readable by "others" are skipped (`permission denied` in the log)

```bash
kubectl -n media exec deploy/navidrome -- touch /music/.rw-test
# touch: /music/.rw-test: Read-only file system
```

## kubernetes (k3s)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: navidrome
  namespace: media
spec:
  replicas: 1
  strategy:
    type: Recreate          # SQLite + RWO PVC: never two pods at once
  selector:
    matchLabels:
      app: navidrome
  template:
    metadata:
      labels:
        app: navidrome
    spec:
      securityContext:
        runAsNonRoot: true
        runAsUser: 65534
        runAsGroup: 65534
      containers:
        - name: navidrome
          image: deluan/navidrome:<pinned-version>
          ports:
            - name: http
              containerPort: 4533
          env:
            - name: ND_MUSICFOLDER
              value: /music
            - name: ND_DATAFOLDER
              value: /data
            - name: ND_SCANSCHEDULE
              value: "1h"
          volumeMounts:
            - name: music
              mountPath: /music
              readOnly: true
            - name: data
              mountPath: /data
          readinessProbe:
            httpGet:
              path: /ping
              port: http
      volumes:
        - name: music
          hostPath:
            path: /path/to/music
            type: Directory   # pod waits if the disk isn't mounted yet
        - name: data
          persistentVolumeClaim:
            claimName: navidrome-data   # local-path, RWO, 2Gi is plenty
```

plus a `LoadBalancer` Service on 4533 (phones don't read `/etc/hosts`,
so they need `http://<server-ip>:4533`) and, optionally, an ingress.

create the admin right after deploying: the first user to log in gets it.

## clients (free)

- Android + iOS: Substreamer (same app on both, from the stores)
- Android: Tempus (open source, active fork of the abandoned Tempo; APK
  from GitHub), Ultrasonic (open source, Play Store / F-Droid)
- iOS: Amperfy (open source)
- desktop: the web UI, or Feishin

check each app's quality setting: "original" means no transcoding.

## check that it streams without transcoding

```bash
kubectl -n media logs deploy/navidrome | grep 'Streaming file'
# format=raw originalFormat=flac originalBitRate=914 -> original file, no transcoding
```

transcoding only happens if a player has it configured (Settings ->
Players) or a client asks for it.

## gotchas

- music on a disk mounted late (encrypted, unlocked at desktop login):
  `hostPath` with `type: Directory` keeps the pod in `ContainerCreating`
  (`FailedMount`) until the folder exists, instead of scanning an empty
  dir. Kubelet retries on its own.
- the running pod keeps the disk busy: `umount` fails with
  `target is busy`. Scale down first:

  ```bash
  kubectl -n media scale deploy/navidrome --replicas=0
  # ... unmount, remount ...
  kubectl -n media scale deploy/navidrome --replicas=1
  ```

- `kubectl apply --dry-run=server` on a file that also creates the
  namespace fails with `namespaces "media" not found`. Expected; the
  real apply works.
- Navidrome shows the tags as they are: `AC/DC`, `AC_DC` and `AC-DC`
  are three artists. Case-only differences (`Iron maiden`) are merged.
  It can't edit tags, by design.
- Songs -> "Shuffle All" queues 500 random songs, not the whole library.
  Press it again for a new batch.
- smart playlists (`.nsp`) are files inside the music folder, so they
  don't work with a read-only library.
