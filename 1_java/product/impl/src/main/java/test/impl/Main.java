package test.impl;

import test.api.Echo;

public class Main {
    public static void main(String[] args) throws ReflectiveOperationException {
        if (args.length != 1) {
            throw new RuntimeException("plugin class must be specified as a first program argument");
        }

        String clazz = args[0];
        Echo echer = loadPlugin(clazz);
        System.out.println("calling plugin: " + clazz);
        System.out.println(echer.echo());
    }

    public static Echo loadPlugin(String className) throws ReflectiveOperationException {
        return (Echo) Class.forName(className).newInstance();
    }
}
