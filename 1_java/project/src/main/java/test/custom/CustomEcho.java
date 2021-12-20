package test.custom;

public class CustomEcho implements test.api.Echo {
    @Override
    public String echo() {
        return "hello from custom code!";
    }
}
