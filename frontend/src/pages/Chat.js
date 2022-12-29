import { useState } from "react"
import { ListContacts, NewContact } from "./Contacts"
import { cookies } from ".."

export const Chat = () => {
    const [status, toggleStatus] = useState(false)
    const AddContact = () => {
        toggleStatus(!status)
    }
    return (
        <div>
            <h1>Hello {cookies.get('firstname')}</h1>
            <h1>This is the Chat Page</h1>
            <h2> Fetching chats...</h2>
            <button onClick={AddContact}>Add Contact</button>
            {status?<NewContact/>:null}
            {/* List out all the contacts (on click, should open a chat) */}
            <ListContacts/>
        </div>
    )
}