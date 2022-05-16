using System.Runtime.InteropServices;

namespace DotnetSample;

public static class MyNativeLibrary
{
    const string C_LIBNAME = @"..\..\..\output\MyNativeLibraryNE.dll";

    [DllImport(C_LIBNAME)] 
    public static extern int MyLibSum(int a, int b);

    [DllImport(C_LIBNAME)] 
    public static extern int MyLibSub(int a, int b);
}