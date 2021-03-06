## tanzu package installed update

Update installed package

```
tanzu package installed update INSTALLED_PACKAGE_NAME [flags]
```

### Examples

```

    # Update installed package with name 'mypkg' with some version to version '3.0.0-rc.1' in specified namespace 	
    tanzu package installed update mypkg --version 3.0.0-rc.1 --namespace test-ns
```

### Options

```
  -h, --help                  help for update
      --install               Install package if the installed package does not exist, optional
  -n, --namespace string      The namespace to locate the installed package which needs to be updated (default "default")
  -p, --package-name string   The public name for the package
  -f, --values-file string    The path to the configuration values file
  -v, --version string        The version which installed package needs to be updated to
```

### Options inherited from parent commands

```
      --kubeconfig string   The path to the kubeconfig file, optional
      --log-file string     Log file path
  -o, --output string       Output format (yaml|json|table)
      --verbose int32       Number for the log level verbosity(0-9)
```

### SEE ALSO

* [tanzu package installed](tanzu_package_installed.md)	 - Manage installed packages

###### Auto generated by spf13/cobra on 8-Jul-2021
