import {useForm} from "react-hook-form"
import { useNavigate } from "react-router-dom";
import * as yup from "yup"
import {yupResolver} from "@hookform/resolvers/yup"
import axios from "axios";
import { cookies, HOSTNAME } from ".."
import { useState } from "react";

export const Register = () => {
    const [status, setStatus] = useState(false);
    let navigate = useNavigate(); 
    const schema = yup.object().shape({
        fname: yup.string().required(),
        lname: yup.string().required(), 
        email: yup.string().email("Provide a valid email").required(),
        pwd: yup.string().min(8, "Atleast 8 characters long").max(20).required(),
        cnfpwd: yup.string().oneOf([yup.ref("pwd")])  
    })
    const {register, handleSubmit, formState : {errors}} = useForm({
        resolver: yupResolver(schema)
    });
    const onSubmit = async (data) => {
        const credentials = { firstname: data.fname, lastname: data.lname, email: data.email, password: data.pwd };
        const res = axios.post(HOSTNAME+'/register', credentials)
            .then(
                (resp) =>{
                    if (resp.data) {
                        // Set the cookie state
                        cookies.set('email', credentials.email, {path: '/'}) // so that its visible in all pages
                        cookies.set('firstname', credentials.firstname, {path: '/'})
                        cookies.set('lastname', credentials.lastname, {path: '/'})
                        navigate("/chat")
                    } else {
                        alert("Account already exists, use a different email address")
                    }
                });
    }

    return (
        <div>
            <h1>This is the Registration Page</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
                <input type="text" placeholder="First Name" {...register("fname")}/> <br/>
                <p>{errors.fname?.message}</p>
                <input type="text" placeholder="Last Name" {...register("lname")}/> <br/>
                <p>{errors.lname?.message}</p>
                <input type="text" placeholder="Email" {...register("email")}/> <br/>
                <p>{errors.email?.message}</p>
                <input type="text" placeholder="Password" {...register("pwd")}/> <br/>
                <p>{errors.pwd?.message}</p>
                <input type="text" placeholder="Confirm Password" {...register("cnfpwd")}/><br/>
                <input type = "submit"/>
            </form>
        </div>
    )
}