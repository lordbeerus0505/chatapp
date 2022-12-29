import { useState, Component, useEffect } from "react"
import { useForm } from "react-hook-form";
import axios from "axios";
import { cookies, HOSTNAME } from ".."

export const NewContact = () => {
    const [email, alterEmail] = useState("");
    const {register, handleSubmit} = useForm()

    const makeRequest = async (data) => {
        console.log(data, "is the data sent to submit, validate in the backend")
        const res = axios.post(HOSTNAME+'/add-contact', {"email": data.mail_id, "firstname": data.firstname, "lastname": data.lastname})
            .then(
                (resp) =>{
                    console.log(resp.data)
                });
    }
    return (
        <div className="new-contact">
            <p>{email}</p>
            <form onSubmit={handleSubmit(makeRequest)}>
                <input onChange={(mail_id) => alterEmail(mail_id.target.value)} placeholder="Enter the email address of the contact" {...register("mail_id")}/><br/>
                <input placeholder="First Name" {...register("f_name")}/><br/>
                <input placeholder="Last Name" {...register("l_name")}/><br/>
                <input type="submit" value="Search"/>
            </form>
        </div>
    )
}

export const ListContacts = () => {
    
    var contacts = [{
        "email":"email1@gmail.com",
        "firstname": "name1",
        "lastname": "name2"
    }]

    useEffect(() => {
        // Fetch the contacts
        retrieveContacts();
    });

    const retrieveContacts = async (data) => {
        console.log(data, "is the data sent to submit, validate in the backend")
        const res = axios.post(HOSTNAME+'/get-chats', {"email": cookies.get("email")}) // create the backend and have it send back from Dynamo all the data.
            .then(
                (resp) =>{
                    console.log(resp.data)
                });
    }
    return (
        <div className="contact-list">
            {
                contacts.map((data) => {
                    return (
                        <div>
                            <h4>{data.firstname} {data.lastname}</h4>
                        </div>
                    )
                })
            }
        </div>
    )
}