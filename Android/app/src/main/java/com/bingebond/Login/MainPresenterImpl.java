package com.bingebond.Login;

import java.util.List;

public class MainPresenterImpl implements MainPresenter {
    MainView mainView;
    //todo - make model :)
    MainModel mainModel;

    public MainPresenterImpl(MainView mainView) {
        this.mainView = mainView;
    }

    @Override
    public void onLoginPressed() {
        mainModel.login(mainView.getLoginparams());
    }

    @Override
    public void onRegisterPressed() {
        List<String> list = mainView.getRegisterparams();
        if(list.get(2).equals(list.get(3))){
            mainModel.register(list);
        }
        else{
            mainView.displaymessage("Passwords don't match");
        }
    }

    @Override
    public void onRegisterpagepressed() {
        mainView.changeLayout();
    }

    @Override
    public void displaySuccess() {
        mainView.displaymessage("Success");
    }

    @Override
    public void displayFailure() {
        mainView.displaymessage("Failed");
    }
}
