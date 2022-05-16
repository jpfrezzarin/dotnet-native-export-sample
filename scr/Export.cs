using System.Runtime.InteropServices;

namespace MyNativeLibrary;

public class Export
{

    /*
     * All exposed methods requires:
     * - Marked with [UnmanagedCallersOnlyAttribute] annotation
     * - Must be public and static
     */

    [UnmanagedCallersOnlyAttribute(EntryPoint = "MyLibSum")]
    public static int MyLibSum(int a, int b) 
    {
        return a + b;
    }

    [UnmanagedCallersOnlyAttribute(EntryPoint = "MyLibSub")]
    public static int MyLibSub(int a, int b) 
    {
        return a - b;
    }
}
