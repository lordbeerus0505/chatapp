import { useState } from "react"
import { ListContacts, NewContact } from "./Contacts"
import { cookies } from ".."

export const Chat = () => {
    const [status, toggleStatus] = useState(false)
    const AddContact = () => {
        toggleStatus(!status)
    }
    return (
        <div className="chat-page">
            <h3 className="header-text">Hello {cookies.get('firstname')}</h3>
            <button className="add-contact" onClick={AddContact}>Add Contact</button>
            <br/><br/>
            {status?<NewContact/>:null}
            {/* List out all the contacts (on click, should open a chat) */}
            <ListContacts/>
        </div>
    )
}