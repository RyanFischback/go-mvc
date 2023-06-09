import React from 'react'
import ReactDOM from "react-dom/client"
import {BrowserRouter, Routes, Route} from 'react-router-dom'
import Home from "../views/home/index";
import About from "../views/about/About";

const root = ReactDOM.createRoot(document.querySelector("#application")!);
function app() {
    root.render(
        <BrowserRouter>
            <Routes>
                <Route index element={<Home />} />
                <Route path="/about" element={<About />} />
            </Routes>
        </BrowserRouter>
    );
}
export default app;