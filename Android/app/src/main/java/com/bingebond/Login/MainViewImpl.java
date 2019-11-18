package com.bingebond.Login;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.Toast;

import com.bingebond.Login.MainView;
import com.bingebond.R;
import com.google.android.material.textfield.TextInputEditText;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

public class MainViewImpl extends AppCompatActivity implements MainView {
    MainPresenter mainPresenter;
    @BindView(R.id.username_input_layout) TextInputEditText Username;
    @BindView(R.id.password_input_layout) TextInputEditText Password;
    @BindView(R.id.register_fullName_input_layout) TextInputEditText Fullname;
    @BindView(R.id.register_email_input_layout) TextInputEditText Email;
    @BindView(R.id.register_password_input_layout) TextInputEditText RegPass;
    @BindView(R.id.register_password_confirm_input_layout) TextInputEditText ConfPass;
    @BindView(R.id.register_contact_input_layout) TextInputEditText Contact;
    @BindView(R.id.Login_button) Button loginbutton;
    @BindView(R.id.go_to_register_button) Button gotoregisterbutton;
    @BindView(R.id.register_button) Button registerbutton;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        ButterKnife.bind(this);
        mainPresenter = new MainPresenterImpl(this);
        loginbutton.setOnClickListener(view -> mainPresenter.onLoginPressed());
        registerbutton.setOnClickListener(view -> mainPresenter.onRegisterPressed());
        gotoregisterbutton.setOnClickListener(view -> mainPresenter.onRegisterpagepressed());
    }
    @Override
    public void displaymessage(String s){
        Toast.makeText(getApplicationContext(),s,Toast.LENGTH_SHORT).show();
    }
    @Override
    public List<String> getRegisterparams(){
        List<String> list = new ArrayList<String>(){{
            add(Fullname.getText().toString());
            add(Email.getText().toString());
            add(RegPass.getText().toString());
            add(ConfPass.getText().toString());
            add(Contact.getText().toString());
        }
        };
        return list;
    }
    @Override
    public List<String> getLoginparams(){
        List<String> list = new ArrayList<String>(){{
            add(Username.getText().toString());
            add(Password.getText().toString());
        }
        };
        return list;
    }

    @Override
    public void changeLayout() {
        setContentView(R.layout.activity_mainreg);
    }


}
