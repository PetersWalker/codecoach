import { Button } from "@mui/material";
import { GoogleLogin } from '@react-oauth/google';
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./App.css";


export default function Login() {
    const [user, setProfile] = useState(null)
    function onSignIn(googleUser) {
        var profile = googleUser.getBasicProfile();

        // console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
        // console.log('Name: ' + profile.getName());
        // console.log('Image URL: ' + profile.getImageUrl());
        // console.log('Email: ' + profile.getEmail()); // This is null if the 'email' scope is not present.
    }

    const navigate = useNavigate();
    return (
        <>
            <GoogleLogin
                onSuccess={credentialResponse => {
                    console.log(credentialResponse);
                    return navigate("/app")
                }}
                onError={() => {
                    console.log('Login Failed');
                }}
            />
            <Button onClick={(e) => {
                navigate("/app")
            }}></Button>
        </>
    )


}