using System.Runtime.InteropServices;

namespace MyNativeLibrary;

public class Export
{

    /*
     * All exposed methods requires:
     * - Marked with [UnmanagedCallersOnlyAttribute] annotation
     * - Must be public and static
     */

    [UnmanagedCallersOnlyAttribute]
    public static int Sum(int a, int b) => a + b;

    [UnmanagedCallersOnlyAttribute]
    public static int Sub(int a, int b) => a - b;
}
