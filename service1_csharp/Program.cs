using System;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Service1CSharp
{
    class Program
    {
        static async Task Main(string[] args)
        {
            // Start the WebSocket server part (handle this in another method or class)

            // Start the client part to connect to the Go server
            using (ClientWebSocket client = new ClientWebSocket())
            {
                Uri serverUri = new Uri("ws://localhost:5001/ws"); // Go server URL
                await client.ConnectAsync(serverUri, CancellationToken.None);
                Console.WriteLine("Connected to the Go server");

                var timer = new Timer(async _ =>
                {
                    var number = new Random().Next(1, 100);
                    Console.WriteLine($"Sending to Go server: {number}");

                    var bytes = Encoding.UTF8.GetBytes(number.ToString());
                    var arraySegment = new ArraySegment<byte>(bytes);
                    await client.SendAsync(arraySegment, WebSocketMessageType.Text, true, CancellationToken.None);
                }, null, TimeSpan.Zero, TimeSpan.FromSeconds(5)); // Adjust the period as needed

                Console.ReadLine(); // Keep the application running
            }
        }
    }
}
