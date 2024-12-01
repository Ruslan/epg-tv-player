import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import App from './App'
import appStore from "./stores/appStore";
import { createContext } from "react";

const container = document.getElementById('root')

const root = createRoot(container)

export const StoreContext = createContext(appStore);

root.render(
    <React.StrictMode>
        <StoreContext.Provider value={appStore}>
            <App />
        </StoreContext.Provider>
    </React.StrictMode>
)
