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
SHA256          E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855       C:\Users\rahul.thakoor\Docume...
```



## References
1. https://en.wikipedia.org/wiki/Hash_function
2. https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/get-filehash?view=powershell-6
3. https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions