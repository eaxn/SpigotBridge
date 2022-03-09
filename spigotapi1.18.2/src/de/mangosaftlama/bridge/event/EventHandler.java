package de.mangosaftlama.bridge.event;

import de.mangosaftlama.bridge.buffer.Buffer;
import de.mangosaftlama.bridge.plugin.SpigotBridge;
import de.mangosaftlama.bridge.server.BridgeServer;
import org.bukkit.event.Listener;
import de.mangosaftlama.bridge.main.Main;
import org.bukkit.event.block.BlockBreakEvent;

import java.io.IOException;
import java.net.DatagramPacket;

public class EventHandler implements Listener {

    @org.bukkit.event.EventHandler
    public void onBlockBreakEvent(BlockBreakEvent e)  {
        if (SpigotBridge.bridgeServer == null) {
            // Wait for connection...
            return;
        }
        BridgeServer.log.info("event (blockbreak)!");
        byte[] Buffer = new byte[2048];
        int bx = e.getBlock().getX();
        int by = e.getBlock().getY();
        int bz = e.getBlock().getZ();
        String x = "New BlockBreakEvent [e] (" + e.getPlayer().getUniqueId().toString() + ") => " + by + "/" + by + "/" + bz;
        Buffer = de.mangosaftlama.bridge.buffer.Buffer.WriteToBuffer(0, Buffer, x.getBytes());
        DatagramPacket packet = new DatagramPacket(Buffer, Buffer.length, SpigotBridge.client, SpigotBridge.port);
        try {
            SpigotBridge.bridgeServer.GetSocket().send(packet);
            BridgeServer.log.info("send event to client!");
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }
}
