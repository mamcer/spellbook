# git

[git commands i run before reading any code](https://piechowski.io/post/git-commands-before-reading-code)

What Changes the Most

```
git log --format=format: --name-only --since="1 year ago" | sort | uniq -c | sort -nr | head -20
```

Who Built This

```
git shortlog -sn --no-merges
```

Where Do Bugs Cluster

```bash
git log -i -E --grep="fix|bug|broken" --name-only --format='' | sort | uniq -c | sort -nr | head -20
```

Is This Project Accelerating or Dying

```bash
git log --format='%ad' --date=format:'%Y-%m' | sort | uniq -c
```

How Often Is the Team Firefighting

```bash
git log --oneline --since="1 year ago" | grep -iE 'revert|hotfix|emergency|rollback'
```

## unstage files

```bash
git reset HEAD [file]  
```

## status

```bash
git status -s  
```

Example:

```bash
M README  
MM Rakefile
A  lib/git.rb  
M  lib/simplegit.rb  
?? LICENSE.txt 
```

## messages

1. La "Trinidad" del día a día (90% de tus commits)
* feat:: Nueva funcionalidad o contenido (ej: una nueva sección en el CV).
* fix:: Corrección de un error (lo que estamos haciendo ahora con el CI).
* docs:: Cambios solo en documentación (README, comentarios, pero no en el CV per se si lo consideras "código").
2. Infraestructura y Calidad
* ci:: Cambios en workflows de GitHub, Jenkins, etc. (Este commit actual encaja perfecto aquí).
* refactor:: Cambios en el código que ni arreglan un bug ni añaden una feature (ej: reordenar el CSS sin cambiar el look).
* chore:: Tareas rutinarias que no afectan el código (actualizar una dependencia, limpiar archivos temporales).

3. El Estilo Profesional (Pragmático 2026)
Un mensaje ideal no se capitaliza y usa el imperativo (como si fuera una orden):

## diff
 
To see what have changed but not yet staged 

```bash
git diff  
```

To see what you have staged and will go in your next commit. --cached is a synonym

```bash
git diff --staged  
```

## removing files

```bash
git rm [file]  
```

remove the file from staging area but keep it in your hard drive

```bash
git rm --cached [file]  
```

## config

Current repository

```bash
git config user.name `<name>`
```

Global

```bash
git config --global user.name `<name>`
git config --global user.email `<email>`
```

Create shortcut to a git command

```bash
git config --global alias `<alias.name>` `<git.command>`
```

Open the global configuration file in a text editor for manual editing

```bash
git config --global --edit
```

## ignoring files

add a .gitignore file  
[https://github.com/github/gitignore](https://github.com/github/gitignore)

## Ignoring Files

```bash
vim .gitignore
```

- Blank lines or lines starting with # are ignored.
- Standard glob patterns work.
- You can start patterns with a forward slash (/) to avoid recursivity
- You can end patterns with a forward slash (/) to specify a directory
- You can negate a pattern by starting it with an exclamation point (!)

## log

```bash
git log
```

Limit the number of commits to `<limit>`

```bash
git log -n `<limit>`  
```

Condensed way. Useful to see a project high level commit history

```bash
git log --oneline  
git log --graph --decorate --oneline
git log --pretty=oneline
git log --pretty=format:"%h - %an, %ar : %s"
```

> *Option	Description of Output*
> 
> %H	Commit hash  
> %h	Abbreviated commit hash  
> %T	Tree hash  
> %t	Abbreviated tree hash  
> %P	Parent hashes  
> %p	Abbreviated parent hashes  
> %an	Author name  
> %ae	Author e-mail  
> %ad	Author date (format respects the --date=option)  
> %ar	Author date, relative  
> %cn	Committer name  
> %ce	Committer email  
> %cd	Committer date  
> %cr	Committer date, relative  
> %s	Subject  

Include which files were altered and the relative number of lines that were added or deleted from each of them

```bash
git log --stat  
```

Display the patch representing each commit. This is the most detailed view

```bash
git log -p 
```

Pattern could be a string or a regular expression

```bash
git log --author="`<pattern>`"  
```

Search for commits with a commit message that matches `<pattern>`. Plain text or a regular expression

```bash
git log --grep="`<pattern>`"  
```

Both commit can be either a commit id  a branch name, HEAD, etc.

```bash
git log `<since>`..`<until>`  
```

History of a particular file

```bash
git log `<file>`  
```

Advanced: commit count - author name - author email 

```bash
git log --format='%an - %ae' | sort | uniq -c | sort -nr
```

## stats

```bash
git shortlog -sn --no-merges
```

#### checkout

Return to master branch. Current state of the project

```bash
git checkout master
```

File `<file>` is changed to the state that it have on `<commit>`

```bash
git checkout `<commit>` `<file>`
```

Update all files in the working folder to match `<commit>`

```bash
git checkout `<commit>`
``` 

Revert any change to `<file>` getting the latest version

```bash
git checkout HEAD `<file>`
```

## revert
 
Generate a new commit that undoes the changes introduced in `<commit>`

```bash
git revert `<commit>`
```

Revert last commit 

```bash
git revert HEAD
```

## reset

Although it’s more often used to undo changes in the staging area and the working directory. In either case, it should only be used to undo local changes—you should never reset snapshots that have been shared with other developers

Remove the specified file from staging area. But leave the working directory unchanged

```bash
git reset `<file>`
```

Reset the staging area to match the most recent commit. But leave the working folder unchanged.

```bash
git reset 
```

Same as git reset but it also tell git to overwrite all changes in local directory

```bash
git reset --hard
```

Move the current branch tip backward to `<commit>`, reset the staging area to match, but leave the working directory alone. All changes made since `<commit>` will reside in the working directory, which lets you re-commit the project history using cleaner, more atomic snapshots

```bash
git reset `<commit>`
```

Move the current branch tip backward to `<commit>` and reset both the staging area and 
the working directory to match. This obliterates not only the uncommitted changes, but all commits after `<commit>`, as well.

```bash
git reset --hard `<commit>`  
```

Remove last two snapshots from repository

```bash
git reset --hard HEAD~2  
```

discard uncommited changes

```bash
git restore .
git clean -fd
```

if you want to discard changes and match the remote

```bash
git fetch origin
git reset --hard origin/dev
git clean -fd
```

### clean

The git clean command removes untracked files from your working directory. Is not undoable

```bash
git clean  
```

Show wich files are going to be removed without actually doing it

```bash
git clean -n  
```

Remove untracked files from the current directory

```bash
git clean -f  
```

Limit the operation to `<path>`

```bash
git clean -f `<path>`  
```

Remove untracked files and folders from the current directory

```bash
git clean -df  
```

Remove untracked files from the current directory as well as files ignored by git

```bash
git clean -xf  
```

## tag

Lightweigh

```bash
git tag [tagname]  
```

Basically the commit checksum stored in a file

Annotated

```bash
git tag -a [tagname] -m [message]  
```

Annotated tags are stored as full objects in the Git database

```bash
git show [tagname]  
```

The tag data along with the commit that was tagged

Show only the files included in commit

```bash
git show --pretty="" --name-only [commit-hash]
```

#### tagging later

```bash
git tag -a [tagname] [checksum]  
```

[checksum] could be part of it. i.e: 9fceb02

#### Sharing Tags

```bash
git push [remote] [tagname]  
```

only tagname is pushed

```bash
git push [remote] --tags  
```

all tags are pushed

#### checking out tags

```bash
git checkout -b [branchname] [tagname]
```

### Unstaging a staged file

```bash
git reset HEAD [file]
```

### branches

List all branches in your repository

```bash
git branch  
```

Create a new branch called `<branch>`. This does not checkout the new branch

```bash
git branch `<branch>`  
```

Deletes branch `<branch>`. Git prevents you from deleteing a  branch if it has unmerged changes

```bash
git branch -d `<branch>`  
```

Force deleted

```bash
git branch -D `<branch>`  
```

Rename the current branch to `<branch>`

```bash
git branch -m `<branch>`  
```

## checkout

Check out the specified branch, which should have already been created with git branch. This makes `<existing-branch>` the current branch, and updates the working directory to match

```bash
git checkout `<existing-match>`  
```

Create and check out new-branch

```bash
git checkout -b `<new-branch>`  
```

Base new branch off of `<existing-branch>` instead current branch

```bash
git checkout -b `<new-branch>` `<existing-branch>`  
```

## merge

Merge the specified branch into the current branch

```bash
git merge `<branch>`  
```

Merge the specified branch into the current branch. But always generate a merge commit. Useful for documenting all merges

```bash
git merge --no--ff `<branch>`  
```

## rewriting git history

Combine the staged changes with the previous commit and replace the previous commit with the resulting snapshot. Running this when there is nothing staged lets you edit the previous commit’s message without altering its snapshot

```bash
git commit --amend  
```

Example

```bash
# Edit hello.py and main.py
git add hello.py
git commit
# Realize you forgot to add the changes from main.py
git add main.py
git commit --amend --no-edit
```

> --no-edit means that there is no change in the commit message

Rebase the current branch onto `<base>`, which can be any kind of commit reference (an ID, a branch name, a tag, or a relative reference to HEAD)

```bash
git rebase `<base>`  
```

Example:

```bash
git checkout new-feature
git rebase master
git checkout master
git merge new-feature
git rebase -i `<base>`
#Interactive rebasing session
```

## remote repositories

List the remote connections you have to other repositories

```bash
git remote   
```

It also shows the url of the remote connections

```bash
git remote -v  
```

Adds a new remote connection

```bash
git remote add `<name>` `<url>`  
```

Removes the `<name>` connection

```bash
git remote rm `<name>`  
```

Renames a remote connection

```bash
git remote rename `<old-name>` `<new-name>`  
```

```bash
git remote show [remote]
```

## fetch  

Fetching is what you do when you want to see what everybody else has been working on. Since fetched content is represented as a remote branch, it has absolutely no effect on your local development work. This makes fetching a safe way to review commits before integrating them with your local repository.

Fetch all of the branches from a repository

```bash
git fetch `<remote>`  
```

Fetch only the branch `<branch>`

```bash
git fetch `<remote>` `<branch>`  
```

## example

```bash
git fetch origin
git checkout master
git log origin/master
git merge origin/master
```

Same as git fetch `<remote>` followed by git merge origin/`<current-branch>`

```bash
git pull `<remote>`  
```

Same as above but instead of using git merge to integrate it uses git rebase

```bash
git pull --rebase `<remote>`  
```

Push the specific branch>` to `<remote>`

```bash
git push `<remote>` `<branch>`  
```

Force the push even if it results in a non-fast forward merge

```bash
git push `<remote>` --force  
```

Push all of your local branches to `<remote>`

```bash
git push `<remote>` --all  
```

Tags are not automatically pushed. The --tags flag send all of your local tags to the remote repository
 
```bash
git push `<remote>` --tags  
```

The following example describes one of the standard methods for publishing local contributions to the central repository

```bash
git checkout master
git fetch origin master
git rebase -i origin/master
# Squash commits, fix up commit messages etc.
git push origin master
```

## remove remote deleted branches locally

Show branches to be deleted

```bash
git remote prune origin --dry-run
```

Locally delete remote deleted branches

```bash
git remote prune origin
```

## user stats

```bash
./user-stats.sh --since="4 week ago"

git log --pretty=format:"%ar - %h - %an: %s"

git log --pretty=format:"%ar - %h - %an: %s" --author='kami'

branches & age
git for-each-ref --sort=committerdate refs/heads/ --format='%(HEAD) %(align:35)%(color:yellow)%(refname:short)%(color:reset)%(end) - %(color:red)%(objectname:short)%(color:reset) - %(align:40)%(contents:subject)%(end) - %(authorname) (%(color:green)%(committerdate:relative)%(color:reset))'
```

another 

```bash
#!/bin/bash
#
# Show user stats (commits, files modified, insertions, deletions, and total
# lines modified) for a repo
git_log_opts=( "$@" )
git log "${git_log_opts[@]}" --format='author: %ae' --numstat \
    | tr '[A-Z]' '[a-z]' \
    | grep -v '^$' \
    | grep -v '^-' \
    | gawk '
        {
            if ($1 == "author:") {
                author = $2;
                commits[author]++;
            } else {
                insertions[author] += $1;
                deletions[author] += $2;
                total[author] += $1 + $2;
                # if this is the first time seeing this file for this
                # author, increment their file count
                author_file = author ":" $3;
                if (!(author_file in seen)) {
                    seen[author_file] = 1;
                    files[author]++;
                }
            }
        }
        END {
            # Print a header
            printf("%-30s\t%-10s\t%-10s\t%-10s\t%-10s\t%-10s\n",
                   "Email", "Commits", "Files",
                   "Insertions", "Deletions", "Total Lines");
            printf("%-30s\t%-10s\t%-10s\t%-10s\t%-10s\t%-10s\n",
                   "-----", "-------", "-----",
                   "----------", "---------", "-----------");
            # Print the stats for each user, sorted by total lines
            n = asorti(total, sorted_emails, "@val_num_desc");
            for (i = 1; i <= n; i++) {
                email = sorted_emails[i];
                printf("%-30s\t%-10s\t%-10s\t%-10s\t%-10s\t%-10s\n",
                       email, commits[email], files[email],
                       insertions[email], deletions[email], total[email]);
            }
        }
'
```
another more

```bash
#!/bin/bash

#for user in "firstname.lastname"

do
  echo 'git-stats -g -s '20 January 2025' --author "'$user'"'
  git-stats -g -s '20 january 2025' --author $user
done

exit 0
```

## ~/.gitconfig

[https://blog.gitbutler.com/how-git-core-devs-configure-git](https://blog.gitbutler.com/how-git-core-devs-configure-git)

```
[column]
        ui = auto
[branch]
        sort = -committerdate
[tag]
        sort = version:refname
[init]
        defaultBranch = main
[diff]
        algorithm = histogram
        colorMoved = plain
        mnemonicPrefix = true
        renames = true
[push]
        default = simple
        autoSetupRemote = true
        followTags = true
[fetch]
        prune = true
        pruneTags = true
        all = true
# why the hell not?
[help]
        autocorrect = prompt
[commit]
        verbose = true
[rerere]
        enabled = true
        autoupdate = true
[core]
        excludesfile = ~/.gitignore
[rebase]
        autoSquash = true
        autoStash = true
        updateRefs = true
# a matter of taste (uncomment if you dare)
[core]
        # fsmonitor = true
        # untrackedCache = true
[merge]
        # (just 'diff3' if git version < 2.3)
        # conflictstyle = zdiff3 
[pull]
        # rebase = true
```

## archivos grandes


```bash
git rev-list --objects --all | grep "$(git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -10 | awk '{print $1}')" 
```

comprimir .git

```bash
git gc --prune=now --aggressive
```

imagenes, usar imagemin


 Para reducir el peso real del repositorio y evitar que Git "recuerde" las versiones pesadas de las imágenes, necesitamos un proceso de dos pasos: Optimización y Reescritura de Historia.

  Aquí tienes el flujo de trabajo profesional para lograrlo:

  1. Optimización con imagemin (o herramientas CLI)

  En lugar de instalar todo el ecosistema de Node para imagemin (que puede ser pesado), como usuario de Linux te recomiendo oxipng para PNGs y jpegoptim para JPGs. Son extremadamente rápidos
  y eficientes.

  Pasos para optimizar:
   1. Instala las herramientas: sudo apt install optipng jpegoptim (en Ubuntu/Debian).
   2. Corre la optimización masiva:

   1    # Para JPGs (los optimiza sin perder calidad perceptible)
   2    find static/img/ -name "*.jpg" -exec jpegoptim --strip-all {} +
   3    # Para PNGs
   4    find static/img/ -name "*.png" -exec optipng -o2 {} +

  2. Cómo eliminar las versiones pesadas de la historia de Git

  Si solo optimizas y haces commit, el archivo .git seguirá pesando lo mismo porque Git guarda la versión original pesada por si alguna vez quieres "volver atrás".

  Para borrar el rastro de esos archivos pesados definitivamente, la herramienta moderna estándar es git-filter-repo (reemplaza al viejo filter-branch).

  El proceso de "limpieza profunda":
   1. Instala la herramienta: sudo apt install git-filter-repo.
   2. Ejecuta la limpieza de archivos grandes (ejemplo para una imagen específica que sepas que pesa mucho):

   1    git filter-repo --path static/img/imagen_pesada.jpg --invert-paths
     Esto borrará esa imagen de TODO el historial de Git como si nunca hubiera existido.

  3. La opción "Pragmática" (Recomendada para un sitio personal)

  Dado que acabas de migrar a Hugo y tienes un nuevo comienzo, si el .git de 57MB te molesta mucho y no necesitas auditar commits de 2017, lo más sano para la performance de tu CI/CD es:

   1. Optimiza las imágenes con los comandos del paso 1.
   2. Haz una copia de tu carpeta mamcer.github.io (solo los archivos, sin el .git).
   3. Borra el repositorio local y vuelve a iniciarlo:

   1    rm -rf .git
   2    git init
   3    git add .
   4    git commit -m "chore: fresh start with optimized assets and Hugo"
   5    git remote add origin https://github.com/mamcer/mamcer.github.io.git
   6    git push -f origin main

  ¿Por qué recomiendo esto?
  Reescribir la historia con filter-repo en un repo con muchos commits pequeños puede ser tedioso. Empezar de cero con los archivos ya optimizados te dejará un .git de apenas 2MB o 3MB,
  haciendo que tus GitHub Actions corran mucho más rápido.

## git diagnostic commands

Before diving into a new codebase, use these Git commands to identify hotspots, bus factors, and project health. This methodology is based on the technical diagnostic approach shared by **Ally Piechowski**.

**Source:** [https://piechowski.io/post/git-commands-before-reading-code/](https://piechowski.io/post/git-commands-before-reading-code/)

---

#### 1. Identify Churn Hotspots
Find the most frequently modified files in the last year. High churn often indicates high complexity or fragile code that requires constant patching.

```bash
git log --format=format: --name-only --since="1 year ago" | sort | uniq -c | sort -nr | head -20
```
*   **Insight:** The top files are likely your biggest "land mines" or areas of technical debt.

#### 2. Bus Factor & Authorship
Identify who built the system and who is currently maintaining it.

```bash
## All-time contributors ranked by commit count
git shortlog -sn --no-merges

## Recent contributors (last 6 months)
git shortlog -sn --no-merges --since="6 months ago"
```
*   **Insight:** If the top all-time contributors aren't in the recent list, you've lost the original context of the codebase.

#### 3. Bug Clusters
Locate where bugs are most frequently reported and fixed.

```bash
git log -i -E --grep="fix|bug|broken" --name-only --format='' | sort | uniq -c | sort -nr | head -20
```
*   **Insight:** Files that appear both in the **Churn Hotspots** and **Bug Clusters** lists are your highest-risk components.

#### 4. Project Momentum
Visualize the project's commit frequency over time to see if development is accelerating, steady, or declining.

```bash
git log --format='%ad' --date=format:'%Y-%m' | sort | uniq -c
```
*   **Insight:** Sudden drops in commits might correlate with team departures, while spikes might indicate "batch" releases instead of continuous delivery.

#### 5. Firefighting Patterns
Measure how often the team is forced to revert or perform emergency hotfixes.

```bash
git log --oneline --since="1 year ago" | grep -iE 'revert|hotfix|emergency|rollback'
```
*   **Insight:** Frequent reverts suggest issues with the deployment pipeline, testing culture, or staging environments.

---

*This summary was generated for the project wiki to establish a standard codebase audit procedure.*