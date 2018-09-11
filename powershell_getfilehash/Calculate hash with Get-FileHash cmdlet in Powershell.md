# Calculate hash with Get-FileHash cmdlet in Powershell

## Introduction
Hashing is the mapping of data of arbitrary size to data of fixed size. The transformation is carried out by passing the data through a hashing algorithm which then generates a unique value representing the content of the data. Hashes are useful in checking data integrity and also computing if two files have the same content. For instance, SUPINFO requires submitting hash values for some projects like Architecture and Development projects submissions. The SCE platform also calculates a hash value when files are uploaded.

In this article, we will describe how to use the `Get-FileHash` cmdlet in powershell to compute the hash value of files.

## Basic Usage

The basic syntax for using the command is as follows:

```
Get-FileHash <path/to/file>
```

For example:

```
Get-FileHash .\example1.txt
```

Which gives the following output:

```
PS C:\Users\rahul.thakoor\Documents> Get-FileHash .\example1.txt

Algorithm       Hash                                                                   Path
---------       ----                                                                   ----
SHA256          10C22747AE63175A5680853FD5180BF758810F2BC122306C06338BCA21B3756B       C:\Users\rahul.thakoor\Docume...
```
By default, powershell displays the output as a table. We can use the `Format-List` cmdlet to format the output as a list of properties in which each property appears on a new line.

```
PS C:\Users\rahul.thakoor\Documents> Get-FileHash .\example1.txt | Format-List


Algorithm : SHA256
Hash      : 10C22747AE63175A5680853FD5180BF758810F2BC122306C06338BCA21B3756B
Path      : C:\Users\rahul.thakoor\Documents\example1.txt

```
We can check if any minor change in the file causes the hash to change:

Check content of original file using the `Get-Content` cmdlet (alias is `cat`):

```
PS C:\Users\rahul.thakoor\Documents> Get-Content .\example1.txt
this is the original document
```

Modify the content of the file by adding another character(`!` here):

```
PS C:\Users\rahul.thakoor\Documents> Get-Content .\example1.txt
this is the original document!
```

Check the hash value again:

```
PS C:\Users\rahul.thakoor\Documents> Get-FileHash .\example1.txt | Format-List


Algorithm : SHA256
Hash      : DD0E86812DCA33D824DF7850802CFECD3B035526DA97991F16085DBA166F21B2
Path      : C:\Users\rahul.thakoor\Documents\example1.txt
```

There is a significant difference in the hash values even when a minor change is made to the content of the file. The value changed from `10C22747AE63175A5680853FD5180BF758810F2BC122306C06338BCA21B3756B` to `DD0E86812DCA33D824DF7850802CFECD3B035526DA97991F16085DBA166F21B2`

As it can be seen from the output, `Get-FileHash` uses the `SHA256` algorithm by default. However, the cmdlet supports changing the algorithm used to calculate the hash.

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
PS C:\Users\rahul.thakoor\Documents> Get-FileHash -Algorithm MD5 .\example1.txt | Format-List


Algorithm : MD5
Hash      : 3E4C5CD8430721A62D370CBFF28CF9AF
Path      : C:\Users\rahul.thakoor\Documents\example1.txt 
```
## References
1. https://en.wikipedia.org/wiki/Hash_function
2. https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/get-filehash?view=powershell-6
3. https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions
4. https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/format-list?view=powershell-6