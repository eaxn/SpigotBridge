package de.mangosaftlama.bridge.cache;

import org.bukkit.event.Event;

public class Cache {

    public static Event[] EventsCache = new Event[]{};

    public static Event GetEvent(int addr) {
        return EventsCache[addr];
    }

    public static void SetEvent(int addr, Event event) {
        EventsCache[addr] = event;
    }
}
