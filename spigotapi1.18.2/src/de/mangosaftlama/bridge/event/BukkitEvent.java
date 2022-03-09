package de.mangosaftlama.bridge.event;

import org.bukkit.event.Event;

public class BukkitEvent {

    private String Type;
    private Event Event;

    public BukkitEvent(Event event, String eventType) {
        Type = eventType;
        Event = event;
    }

    public String getType() {
        return Type;
    }

    public void setType(String type) {
        Type = type;
    }

    public org.bukkit.event.Event getEvent() {
        return Event;
    }

    public void setEvent(org.bukkit.event.Event event) {
        Event = event;
    }
}
