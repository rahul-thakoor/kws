# Calculate hash with Get-FileHash cmdlet in Powershell

## Introduction
Hashing is the mapping of data of arbitrary size to data of fixed size. The transformation is carried out by passing the data through a hashing algorithm which then generates a unique value representing the content of the data. Hashes are useful in checking data integrity and also computing if two files have the same content. For instance, SUPINFO requires submitting hash values for projects submissions of several modules like Architecture and Development . The SCE platform also calculates a hash value when files are uploaded.

In this article, we will describe how to use the `Get-FileHash` cmdlet in powershell to compute the hash value of files.

## Basic Usage

The basic syntax for using the command is as follows:

```
Get-FileHash <path\to\file>
```

For example:

```
Get-FileHash .\example1.txt
```

Which gives the following output:

```
PS C:\Users\supinfo\Documents> Get-FileHash .\example1.txt

Algorithm       Hash                                                                   Path
---------       ----                                                                   ----
SHA256          10C22747AE63175A5680853FD5180BF758810F2BC122306C06338BCA21B3756B       C:\Users\supinfo\Docume...
```
By default, powershell displays the output as a table. We can use the `Format-List` cmdlet to format the output as a list of properties in which each property appears on a new line.

```
PS C:\Users\supinfo\Documents> Get-FileHash .\example1.txt | Format-List


Algorithm : SHA256
Hash      : 10C22747AE63175A5680853FD5180BF758810F2BC122306C06338BCA21B3756B
Path      : C:\Users\supinfo\Documents\example1.txt

```
We can check if any minor change in the file causes the hash to change:

Check content of original file using the `Get-Content` cmdlet (alias is `cat`):

```
PS C:\Users\supinfo\Documents> Get-Content .\example1.txt
this is the original document
```

Modify the content of the file by adding another character(`!` here):

```
PS C:\Users\supinfo\Documents> Get-Content .\example1.txt
this is the original document!
```

Check the hash value again:

```
PS C:\Users\supinfo\Documents> Get-FileHash .\example1.txt | Format-List


Algorithm : SHA256
Hash      : DD0E86812DCA33D824DF7850802CFECD3B035526DA97991F16085DBA166F21B2
Path      : C:\Users\supinfo\Documents\example1.txt
```

There is a significant difference in the hash values even when a minor change is made to the content of the file. The value changed from `10C22747AE63175A5680853FD5180BF758810F2BC122306C06338BCA21B3756B` to `DD0E86812DCA33D824DF7850802CFECD3B035526DA97991F16085DBA166F21B2`

Alternatively, we can check if different files have the same content. If two files have identical hash values, the contents of the files are also identical.

For example, let's assume we have the following two files:

```
    Directory: C:\Users\supinfo\Documents\files\duplicate


Mode                LastWriteTime         Length Name
----                -------------         ------ ----
-a----        9/12/2018   8:39 AM            358 2NET.cs
-a----        9/12/2018   8:38 AM            358 2NET_backup.cs

```

Each have the same content:

```
PS C:\Users\supinfo\Documents\files\duplicate> Get-Content .\2NET.cs
```
```
// A Hello World! program in C#.
using System;
namespace HelloWorld
{
    class Hello
    {
        static void Main()
        {
            Console.WriteLine("Hello World!");

            // Keep the console window open in debug mode.
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }
    }
}
```

```
PS C:\Users\supinfo\Documents\files\duplicate> Get-Content .\2NET_backup.cs
```
```
// A Hello World! program in C#.
using System;
namespace HelloWorld
{
    class Hello
    {
        static void Main()
        {
            Console.WriteLine("Hello World!");

            // Keep the console window open in debug mode.
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }
    }
}
```

```
Algorithm : SHA256
Hash      : 827167A7955A3CAD59D30DB0466873D9E78CC18F563DE02D7C514F9304A25326
Path      : C:\Users\supinfo\Documents\files\duplicate\2NET.cs

Algorithm : SHA256
Hash      : 827167A7955A3CAD59D30DB0466873D9E78CC18F563DE02D7C514F9304A25326
Path      : C:\Users\supinfo\Documents\files\duplicate\2NET_backup.cs

```
The hash value does not depend on the filename and file extension. It only depends on the content of the file.

## Specifying the algorithm to use for computing hash values
The `-Algorithm` switch allows the user to specify the cryptographic hash function to use for computing the hash value. As used above, if no option is supplied, the cmdlet uses the `SHA256` algorithm by default. The acceptable parameters are:
- SHA1
- SHA256
- SHA384
- SHA512
- MD5

E.g:

Computing the `MD5 checksum` of a file:

```
PS C:\Users\supinfo\Documents> Get-FileHash -Algorithm MD5 .\example1.txt | Format-List


Algorithm : MD5
Hash      : 3E4C5CD8430721A62D370CBFF28CF9AF
Path      : C:\Users\supinfo\Documents\example1.txt 
```

## Computing the hash values for several files

We can feed a string containing path to several files and compute the hash values of the specified files.

For instance, let's assume we have the following files:

```console
PS C:\Users\supinfo\Documents\files> ls


    Directory: C:\Users\supinfo\Documents\files


Mode                LastWriteTime         Length Name
----                -------------         ------ ----
-a----        9/12/2018   8:16 AM             37 1ADS-MP.py
-a----        9/12/2018   8:16 AM             37 2ADS-MP.py
-a----        9/12/2018   8:17 AM           1236 2CMP-compiler.yy
-a----        9/12/2018   8:17 AM             10 2LIN-script.sh
-a----        9/12/2018   8:18 AM             29 2WEB-Assignment.js
-a----        9/12/2018   8:19 AM             65 KWS-todo.txt
```

To calculate the hash values of individual files, use the following commands:

```
Get-ChildItem | Get-FileHash | Format-List
```

`Get-ChildItem` lists contents in the folder which is piped into the `Get-FileHash` cmdlet.

```
PS C:\Users\supinfo\Documents\files> Get-ChildItem | Get-FileHash | Format-List


Algorithm : SHA256
Hash      : C96B7E3F9A89CC21D6722658781670083205145CC3FFBED78B55EDFD05FBBEC2
Path      : C:\Users\supinfo\Documents\files\1ADS-MP.py

Algorithm : SHA256
Hash      : D615694B152F5BB7180509EAFA2E17456F5EE34B2F7BB6D7B29BB383B59C449D
Path      : C:\Users\supinfo\Documents\files\2ADS-MP.py

Algorithm : SHA256
Hash      : 613367E008E5DD74D01CEE9DEDF82CF72FA86CB5BA1619852CA4F25689EC700C
Path      : C:\Users\supinfo\Documents\files\2CMP-compiler.yy

Algorithm : SHA256
Hash      : 9C58C2AFB8A292AFBD285F0EF179E8C65E83AEEDF007A3C2728664649B1A0A30
Path      : C:\Users\supinfo\Documents\files\2LIN-script.sh

Algorithm : SHA256
Hash      : 11D289CEA607D6C371CCBE9C94CC74AD2EF8F864CDC93FCCD0B8C713638707DC
Path      : C:\Users\supinfo\Documents\files\2WEB-Assignment.js

Algorithm : SHA256
Hash      : 5F193E7C3AFC1B13D6C3AFF2341E049AE731282545EA76EB28EF62651A28C50C
Path      : C:\Users\supinfo\Documents\files\KWS-todo.txt
```

Ofcourse, we can write the output to a file using `Out-File`:

```
Get-ChildItem | Get-FileHash | Format-List | Out-File -FilePath ../hashvalues.txt
```

## Conclusion

This article introduces the `Get-FileHash` cmdlet which can be used to compute hash values of specified files. The cmdlet can be used in conjunction with other cmdlets to achieve more productive results. `Get-Filehash` also allows specifying which algorithm to use for computing hash values.

## References
1. https://en.wikipedia.org/wiki/Hash_function
2. https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/get-filehash?view=powershell-6
3. https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions
4. https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/format-list?view=powershell-6
5. https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/out-file?view=powershell-6
6. https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/inside-a-program/hello-world-your-first-program