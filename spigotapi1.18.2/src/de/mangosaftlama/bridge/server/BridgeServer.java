package de.mangosaftlama.bridge.server;

import de.mangosaftlama.bridge.env.Environment;
import de.mangosaftlama.bridge.main.Main;
import de.mangosaftlama.bridge.plugin.SpigotBridge;

import java.io.IOException;
import java.net.*;
import java.util.logging.Logger;


public class BridgeServer extends Thread{

    private DatagramSocket socket;
    private boolean running;
    private byte[] buf = new byte[2048];

    public static Logger log = Logger.getLogger("s");

    int port = 13478;

    public DatagramSocket GetSocket() {
        return socket;
    }

    public BridgeServer() throws SocketException {
        log.info("created new server");
        boolean exception = true;
        while(exception){
            try {
                socket = new DatagramSocket(port);
                exception = false;
            } catch (BindException e) {
                port++;
                continue;
            }
            break;
        }
        log.info("server is on " + port + ".");

    }

    public void run() {
        buf = new byte[2048];
        running = true;

        while (running) {
            DatagramPacket packet
                    = new DatagramPacket(buf, buf.length);
            try {
                socket.receive(packet);
            } catch (IOException e) {
                e.printStackTrace();
            }

            InetAddress address = packet.getAddress();
            int port = packet.getPort();
            packet = new DatagramPacket(buf, buf.length, address, port);
            String received
                    = new String(packet.getData(), 0, packet.getLength());

            if (!Environment.allowNonLocalhost && !address.getHostAddress().equals("127.0.0.1")) {
                log.warning("Unknown address (" + address.getHostAddress() + ") tries to connect.");
                continue;
            }
            SpigotBridge.client = address;
            SpigotBridge.port = packet.getPort();

            if (received.contains("ping!")) {
                log.info("received [ping!] from client (" + address.getHostAddress() + ").");
                byte[] Buffer = new byte[2048];
                Buffer[0] = 'o';
                Buffer[1] = 'k';
                Buffer[2] = '!';
                log.info("sending [ok!] to client (" + address.getHostAddress() + ").");
                DatagramPacket x = new DatagramPacket(Buffer, buf.length, address, port);
                try {
                    socket.send(x);
                } catch (IOException e) {
                    e.printStackTrace();
                }
                continue;
            } else if (received.contains("im hearing you to XD!")) {
                log.info("starting to listen to client!");
                byte[] Buffer = new byte[2048];
                Buffer[0] = 'o';
                Buffer[1] = 'k';
                Buffer[2] = '!';
                log.info("sending [ok!] for hearing to client (" + address.getHostAddress() + ").");
                DatagramPacket x = new DatagramPacket(Buffer, buf.length, address, port);
                try {
                    socket.send(x);
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            /*
            try {
                socket.send(packet);
            } catch (IOException e) {
                e.printStackTrace();
            }
             */
        }
        socket.close();
    }
}
