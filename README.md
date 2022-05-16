# .NET Native Export Sample

A .NET project that uses the DNNE to export it as native library, to be consumed for other applications even developed in another language.

Based on [AaronRobinsonMSFT/DNNE](https://github.com/AaronRobinsonMSFT/DNNE)


## Stack

- .NET 6.0
- Phyton 3.10 (for sample)
- Go 1.18.2 (for sample)


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

## Test Samples

To test the present samples, first, builds the `scr/MyNativeLibrary.csproj`.

```
dotnet build scr/MyNativeLibrary.csproj -o output
```

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

To test the Python sample, run the command bellow:

```
python samples/python/run.py
```

### Test Go Sample

To test the Go sample, run the command bellow:

```
go run samples/go/run.go
```