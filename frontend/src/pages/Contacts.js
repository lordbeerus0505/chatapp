import { useState, Component, useEffect } from "react"
import { useForm } from "react-hook-form";
import axios from "axios";
import { cookies, HOSTNAME } from ".."
import "../App.css"

export const NewContact = () => {
    const [email, alterEmail] = useState("");
    const {register, handleSubmit} = useForm()

    const makeRequest = async (data) => {
        console.log(data, "is the data sent to submit, validate in the backend")
        const res = axios.post(HOSTNAME+'/add-contact', {"email": data.mail_id, "firstname": data.f_name, "lastname": data.l_name, "user": cookies.get("email")})
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
    
    var [contacts, setContacts] = useState([])

    useEffect(() => {
        // Fetch the contacts
        retrieveContacts();
    }, []);

    const retrieveContacts = async (data) => {
        console.log(data, "is the data sent to submit, validate in the backend")
        const res = axios.post(HOSTNAME+'/get-chats', {
            "email": cookies.get("email"),
            "firstname": cookies.get("firstname"),
            "lastname": cookies.get("lastname")
        }).then(
                (resp) =>{
                    if (resp.email != "False") {
                        
                        setContacts(resp.data.Contacts)
                    }
                    console.log(contacts)
                });
    }

    const retrieveChat = (email) => {
        console.log(email, 'is the email PK')
        // Send request to backend to retrieve the conversation between current user (in cookie) and email user
    }
    return (
        <div className="chat-window">
            <div className="contact-list">
                {
                    contacts.map((data) => {
                        return (
                            <div >
                                <button className="chat-names" onClick={() => retrieveChat(data.Email)}>{data.FirstName} {data.LastName}</button>
                            </div>
                        )
                    })
                }
            </div>
        </div>
        
    )
}