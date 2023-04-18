using System;

public class ConvertTemp
{
    public static double convert(double temp)
    {
        
        return ((temp-32)*5)/9;
    }
    public static void Main()
    {
        Console.WriteLine("Enter temperature in Farenheit");
        string tempF = Console.ReadLine();
        double doubletempF = Convert.ToDouble(tempF);
        // Console.WriteLine(tempF);
        Console.WriteLine (convert(doubletempF));
    }
}