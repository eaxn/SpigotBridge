package de.mangosaftlama.bridge.buffer;

public class Buffer {

    public static byte[] WriteToBuffer(int s, byte[] buffer, byte[] data) {
        for (int x = 0; x < data.length; x++) {
            buffer[s+x] = data[x];
        }
        return data;
    }
}
