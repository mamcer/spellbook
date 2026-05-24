# dotnet 

## useful commands

```bash
dotnet --info
```

```bash
dotnet new list
```

```bash
dotnet build --configuration Relase
```

```bash
dotnet test --filter Category=unit
```

```bash
dotnet add package Serilog
```

```bash
dotnet list package --outdated 
```

```bash
dotnet tool install -g dotnet-ef
dotnet tool restore
```

```bash
dotnet watch test
```

```bash
dotnet --list-sdks
dotnet --list-runtimes
```

```bash
# on project directory
dotnet publish -c Release -r linux-x64 /p:PublishSingleFile=true --self-contained true
```
> - by default it is stored in `bin/Relase/net8.0/linux-x64/publish`  
> - you can add `-o` option to configure output directory, for example `-o publish/` 
> - optionally you can add `/p:PublishTrimmed=true` to reduce binary size (this is not recommended if for example you use entity framework core)  

## logging

```csharp
using Serilog;

Log.Error("x-signature, [{x-signature}]", sig);
Log.Error("x-request-id, [{x-request-id}]", reqId);
```

## net core packages

[https://github.com/efcore/EFCore.NamingConventions](https://github.com/efcore/EFCore.NamingConventions)

## issues

### System.IO.IOException: The configured user limit (128) on the number of inotify instances has been reached

On the root directory add: 

`.watchconfig`

```json
{
  "watch": [
    {
      "include": [
        "**/*.cs",
        "**/*.csproj",
        "**/*.json",
        "**/*.cshtml",
        "**/*.razor"
      ]
    },
    {
      "exclude": [
        "**/bin/**",
        "**/obj/**",
        "**/node_modules/**",
        "**/.git/**",
        "**/wwwroot/**",
        "**/logs/**",
        "**/.vs/**",
        "**/.vscode/**"
      ]
    }
  ]
}
```

run:

```bash
#!/bin/bash
# Script to increase inotify limits on Linux
# This fixes the "configured user limit (128) on the number of inotify instances has been reached" error

# Temporary increase (current session only)
echo "Setting temporary inotify limits for current session..."
sudo sysctl fs.inotify.max_user_instances=512
sudo sysctl fs.inotify.max_user_watches=524288

# Permanent increase (survives reboot)
echo "Setting permanent inotify limits..."
echo "fs.inotify.max_user_instances=512" | sudo tee -a /etc/sysctl.conf
echo "fs.inotify.max_user_watches=524288" | sudo tee -a /etc/sysctl.conf

echo "Done! The limits have been increased."
echo "Current limits:"
sysctl fs.inotify.max_user_instances
sysctl fs.inotify.max_user_watches
echo ""
echo "Note: You may need to restart your application or IDE for changes to take effect."
```