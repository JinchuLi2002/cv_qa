using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;

namespace Service1CSharp
{
    class Program
    {
        static async Task Main(string[] args)
        {
            // Start the client part to send an HTTP POST request to the Go server
            using (HttpClient httpClient = new HttpClient())
            {
                Uri serverUri = new Uri("http://go-server:5001/api/send"); // Updated URL
                Console.WriteLine("Sending HTTP POST request to the Go server");

                while (true) // Run indefinitely
                {
                    var number = new Random().Next(1, 100);
                    Console.WriteLine($"Sending to Go server: {number}");

                    var data = new
                    {
                        Data = number // Match the field name in your Go struct
                    };

                    var jsonString = JsonSerializer.Serialize(data);
                    var content = new StringContent(jsonString, Encoding.UTF8, "application/json");

                    // Send the HTTP POST request
                    var response = await httpClient.PostAsync(serverUri, content);

                    if (response.IsSuccessStatusCode)
                    {
                        // Read and display the response if needed
                        var responseContent = await response.Content.ReadAsStringAsync();
                        Console.WriteLine($"Response from Go server: {responseContent}");
                    }
                    else
                    {
                        Console.WriteLine($"HTTP request failed with status code: {response.StatusCode}");
                    }

                    // Delay for 5 seconds before sending the next request
                    await Task.Delay(TimeSpan.FromSeconds(5));
                }
            }
        }
    }
}
