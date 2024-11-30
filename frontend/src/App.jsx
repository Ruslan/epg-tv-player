import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Channel from "./pages/Channel";
import Videos from "./pages/Videos";
import Header from "./components/Header";
import Settings from "./pages/Settings";

function App() {
    return (
        <Router>
            <Header />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/channel/:slug" element={<Channel />} />
                <Route path="/videos" element={<Videos />} />
                <Route path="/settings" element={<Settings />} />
            </Routes>
        </Router>
    );
}

export default App;
