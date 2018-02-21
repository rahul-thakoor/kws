# Powershell equivalents for some common unix commands

## Introduction

Powershell [1] is a task-based command-line tool and a scripting language. Originally developed for Windows, on 18 August 2016 Microsoft open sourced Powershell with support for other platforms such as macOS and Linux. Powershell uses the "Verb-Noun" syntax for commands which are known as _cmdlets_. In this article we are going to introduce some powershell cmdlets for common unix commands.

## Getting help
On unix systems, the command `man` gives help on a specific command. Powershell provides information about commands and concepts through the `Get-Help` cmdlet [2].

```shell
Get-Help <cmdlet>
```

e.g 
```shell
    Get-Help Get-Location

NAME
    Get-Location
    
SYNTAX
    Get-Location [-PSProvider <string[]>] [-PSDrive <string[]>]  [<CommonParameters
    >]
    
    Get-Location [-Stack] [-StackName <string[]>]  [<CommonParameters>]
    

ALIASES
    gl
    pwd

``` 

To list all help topics on the system:

```shell
    Get-Help *
```

Powershell also provides contextual help on topics. Use

```
Get-Help about_\<topic-name\>
```
e.g: Getting help on _Comparison_Operators_ :

```
Get-Help about_Comparison_Operators
```


## Display working directory name

On unix systems, `pwd` returns the current working directory name. The equivalent for Powershell is `Get-Location` [3] cmdlet. 

```shell
PS /bin> Get-Location                                                               

Path
----
/bin

```
Fortunately, powershell has built-in aliases for common unix commands. Therefore,

```
    pwd
```
is same as running the `Get-Location` cmdlet.

```
PS /bin> pwd                                                                        

Path
----
/bin

```

## 


## References
- [1] https://docs.microsoft.com/en-us/powershell/scripting/powershell-scripting?view=powershell-6

- [2] https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/get-help?view=powershell-6

- [3] https://docs.microsoft.com/en-us/powershell/module/Microsoft.PowerShell.Management/Get-Location?view=powershell-6