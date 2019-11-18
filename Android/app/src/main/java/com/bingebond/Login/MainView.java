package com.bingebond.Login;


import java.util.List;

public interface MainView {
    void displaymessage(String s);
    List<String> getRegisterparams();
    List<String> getLoginparams();
    void changeLayout();
}
