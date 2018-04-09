# Building A .NET Core Console App for Raspberry Pi

## Introduction

.NET Core is an opensource development platform maintained by Microsoft and the .NET community.[1] The current stable release is .NET Core 2.0. .NET Core consists of the .NET runtime, framework libraries and .NET Core SDK. .NET Core is cross -platform and runs on Windows, macOS and GNU/Linux however it currently only supports x86 and x64 platforms.[2] As of writing, there is no SDK that runs on ARM32 architecture which is usually found in Raspberry Pis.[3] 
However, there are ARM32 builds which allow developing .NET Core applications on x86/x64 hosts, then publishing the application to run on Raspberry Pi with ARMv7 or ARMv8 CPU. 

In this article, we are going to demonstrate how to develop a basic .NET Core console application on an x86/x64 host running 
either Windows, macOS or GNU/Linux. Then, we are going to publish the application to run on Raspberry Pi.
## References

- [1] https://docs.microsoft.com/en-us/dotnet/core/
- [2] https://github.com/dotnet/core/blob/master/release-notes/2.0/2.0-supported-os.md
- [3] https://github.com/dotnet/core/blob/master/samples/RaspberryPiInstructions.md