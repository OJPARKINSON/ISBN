// See https://aka.ms/new-console-template for more information
using System.Runtime.InteropServices;
using System.Text.RegularExpressions;


class ISBN
{
  public string? Isbn { get; set; }

  public void validate()
  {
    string pattern = @"^(?=(?:\D*\d){10}(?:(?:\D*\d){3})?$)[\d-]+$";
    // Create a Regex
    Regex rg = new Regex(pattern);

    if (!rg.IsMatch(this.Isbn))
    {
      Console.WriteLine("ISBN failed validation");
    }
  }
}

class Program
{
  static void Main(string[] args)
  {
    ISBN isbn = new ISBN();
    isbn.Isbn = "0-596-52068-1";

    isbn.validate();
  }
}
