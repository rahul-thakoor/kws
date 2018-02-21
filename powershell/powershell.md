# Powershell equivalents for some common unix commands

![powershell](PowerShell.png)

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
    Get-Location [-PSProvider <string[]>] [-PSDrive <string[]>]  [<CommonParameters>]
    
    Get-Location [-Stack] [-StackName <string[]>]  [<CommonParameters>]
    
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

## Changing Directory

`cd` allows a unix user to set the current working location to a specified location.. The powershell equivalent is `Set-Location`[4] with aliases `sl`, `cd` and `chdir`.

## Listing Directory Contents
On unix, the command `ls` displays contents of a specified location. Powershell has a similar alias, which references the `Get-ChildItem` cmdlet. Similar to unix, we can add options to a particular cmdlet.
For e.g, listing only files with `.txt` extension:

```shell
Get-ChildItem -Path *.txt                                          


    Directory: /Users/supinfo


Mode                LastWriteTime         Length Name                              
----                -------------         ------ ----                              
------       03/21/2017     11:39             25 offices.txt                       
------       03/21/2017     11:41             46 phones.txt                        
------       04/21/2017     08:39              0 untitled.txt
```

## Creating new items
On unix systems, `touch` and `mkdir` allow us to create a new file or new directory respectively. Powershell cmdlet `New-Item` [5] can create both files and directories. We need to use the _ItemType_ parameter, otherwise the default creates a new file.
 Creating a file:
```shell
New-Item -Name "newFile.txt" -ItemType File                                                                            


    Directory: /Users/supinfo/Desktop


Mode                LastWriteTime         Length Name                                                                                          
----                -------------         ------ ----                                                                                          
------       02/21/2018     14:51              0 newFile.txt
```

Creating a directory:

```shell
New-Item -Name "newDir" -ItemType Directory                                                                            


    Directory: /Users/supinfo/Desktop


Mode                LastWriteTime         Length Name                                                                                          
----                -------------         ------ ----                                                                                          
d-----       02/21/2018     14:52                newDir
```

`New-Item` can also create symbolic or hard links. On unix systems, this requires the `ln` command.

```shell
New-Item -Name link -ItemType SymbolicLink -Value ./newFile.txt                                                        


    Directory: /Users/supinfo/Desktop


Mode                LastWriteTime         Length Name                                                                                          
----                -------------         ------ ----                                                                                          
-----l       02/21/2018     15:11             32 link  

```

## More commands to explore
Powershell has numerous cmdlets and utilities. To get a quick reference to some common commands, look for commands with aliases. 

```shell
Get-Alias *                                                                                                            

CommandType     Name                                   Version    Source                                               
-----------     ----                                   -------    ------                                               
Alias           ? -> Where-Object                                                                                      
Alias           % -> ForEach-Object                                                                                    
Alias           cd -> Set-Location                                                                                     
Alias           chdir -> Set-Location                                                                                  
Alias           clc -> Clear-Content                                                                                   
Alias           clear -> Clear-Host                                                                                    
Alias           clhy -> Clear-History  
```

## References
- [1] https://docs.microsoft.com/en-us/powershell/scripting/powershell-scripting?view=powershell-6

- [2] https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/get-help?view=powershell-6

- [3] https://docs.microsoft.com/en-us/powershell/module/Microsoft.PowerShell.Management/Get-Location?view=powershell-6

- [4] https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.management/set-location?view=powershell-6

- [5] https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.management/new-item?view=powershell-6