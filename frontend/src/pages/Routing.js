import { Routes, Route, Link } from "react-router-dom"
import {Home, Chat, Login, Register } from "../index.js"
export const RoutingPages = () => {
    return (
        <Routes>
          <Route path="/" element={<Home/>}/>
          <Route path="/login" element={<Login/>}/>
          <Route path="/register" element={<Register/>}/>
          <Route path="/chat" element={<Chat/>}/>
        </Routes>
    )
}

export const HomePageLinks = () => {
    return (
        <div>
            <Link to="/">Home</Link>
            <Link to="/login">Login</Link>
            <Link to="/chat">Chat</Link>
        </div>
    )
        
}