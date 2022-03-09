package de.mangosaftlama.bridge.plugin;

import de.mangosaftlama.bridge.event.EventHandler;
import de.mangosaftlama.bridge.server.BridgeServer;
import org.bukkit.Bukkit;
import org.bukkit.plugin.java.JavaPlugin;

import java.net.InetAddress;
import java.net.SocketException;

public class SpigotBridge extends JavaPlugin {


    public static BridgeServer bridgeServer;
    public static InetAddress client;
    public static int port;




    @Override
    public void onEnable() {
        try {
            bridgeServer = new BridgeServer();
        } catch (SocketException e) {
            e.printStackTrace();
        }
        bridgeServer.start();
        BridgeServer.log.info("Enabled the plugin.");
        BridgeServer.log.info("Starting the server...");
        Bukkit.getPluginManager().registerEvents(new EventHandler(), this);
    }

    @Override
    public void onDisable() {
        BridgeServer.log.info("Disabled the plugin.");
    }
}
