import {useForm} from "react-hook-form"
import * as yup from "yup"
import {yupResolver} from "@hookform/resolvers/yup"
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { HOSTNAME, cookies } from "..";
import { useState } from "react";



export const Login = () => {
    const [status, setStatus] = useState(false)
    let navigate = useNavigate(); 
    const schema = yup.object().shape({
        email: yup.string().email('Invalid Email Address').required(),
        pwd: yup.string().min(4, 'Password must be atleast 4 characters long').max(20).required(),
    })
    const {register, handleSubmit, formState: {errors}} = useForm({
        resolver: yupResolver(schema)
    });
    const onSubmit = async (data) => {
        // Send data to go backend

        const credentials = { email: data.email, password: data.pwd };
        const res = await axios.post(HOSTNAME+'/login', credentials)
            .then(
                (resp) =>{
                    console.log(resp.data)
                    if (resp.data.Email) {
                        // Set the cookie state
                        cookies.set('email', resp.data.Email, {path: '/'}) // so that its visible in all pages
                        cookies.set('firstname', resp.data.FirstName, {path: '/'})
                        cookies.set('lastname', resp.data.LastName, {path: '/'})
                        navigate("/chat")
                    } else {
                        alert("Invalid credentials, try again!")
                    }
                });        
    }

    const ForgotPassword = () => {
        // ToDo. Add support for forgot password
        alert("Haha! Sucks to be you.")
    }

    const CreateAccount = () => {
        let path = `/register`; 
        navigate(path);
    }
    return (
        <div>
            <h1>This is the Login Page</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
                <input type="text" placeholder="Email Address" {...register("email")}/> <br/>
                <p>{errors.email?.message}</p>
                <input type="text" placeholder="Password" {...register("pwd")}/> <br/>
                <p>{errors.pwd?.message}</p>
                <input type="submit" value="Login"/>
                
            </form>
            <button onClick={ForgotPassword}>Forgot Password</button> <br/>
            <button onClick={CreateAccount}>Create an Account</button>
        </div>
    )
}