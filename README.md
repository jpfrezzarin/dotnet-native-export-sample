# .NET Native Export Sample

A .NET project that uses the DNNE to export it as native library, to be consumed for other applications even developed in another language.

Based on [AaronRobinsonMSFT/DNNE](https://github.com/AaronRobinsonMSFT/DNNE)


## Stack

- .NET 6.0
- Phyton 3.10 (for sample)
- Go 1.18.2 (for sample)
- [clang](https://clang.llvm.org/) (for builds in linux and macOS)


## How it works

The `.csproj` of .NET class library, which going to generate the native library, should contains these lines below:

```
<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <TargetFramework>net6.0</TargetFramework>
    <ImplicitUsings>enable</ImplicitUsings>
    <Nullable>enable</Nullable>
    <EnableDynamicLoading>true</EnableDynamicLoading> <!-- ADD THIS --> 
    <DnneAddGeneratedBinaryToProject>true</DnneAddGeneratedBinaryToProject> <!-- ADD THIS --> 
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="DNNE" Version="1.*" /> <!-- ADD THIS --> 
  </ItemGroup>

</Project>
```

On the method that is going to be exposed, mark it as `public` and `static`, with `[UnmanagedCallersOnlyAttribute]` annotation, like the method below:

```
[UnmanagedCallersOnlyAttribute]
public static int Sum(int a, int b) 
{
    return a + b;
}
```

## Build

To build the solution, run the command bellow in terminal:

```
dotnet build scr/MyNativeLibrary.csproj -o output
```

## Publish

To build the solution, run the command bellow in terminal:

```
dotnet publish scr/MyNativeLibrary.csproj -o output -c Release -r {SID} --self-contained
```

Where SID is dotnet runtime (e.g.: `win-x64`, `linux-x64`)

## Test

> To test the samples, first, builds or publish the solution

### Test .NET Sample

To test the .NET sample, builds the sample:

```
dotnet build samples/dotnet/DotnetSample.csproj -o samples/dotnet
```

After, run the command bellow:

```
dotnet samples/dotnet/output/DotnetSample.dll
```

### Test Python Sample

To test the Python sample in Windows,, run the command bellow:

```
python samples/python/windows.py
```

To test the Python sample in a Linux distribution, run the command bellow:

```
python3 run samples/python/linux.py
```

### Test Go Sample

To test the Go sample in Windows, run the command bellow:

```
go run samples/go/windows.go
```

To test the Go sample in a Linux distribution, run the command bellow:

```
go run samples/go/linux.go
```
