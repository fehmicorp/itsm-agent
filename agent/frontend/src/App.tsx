import {useEffect, useState} from 'react';
import { EventsOn } from '../wailsjs/runtime/runtime'
import Dashboard from './pages/Dashboard'
import Scan from './pages/Scan'
import Connection from './pages/Connection'
import { HashRouter, Route, Routes, useNavigate } from 'react-router-dom';

function NavigationHandler() {
    const navigate = useNavigate();
    useEffect(() => {
        const unlisten = EventsOn("navigate", (path: string) => {
            if (path) {
                navigate(path);
            }
        });
        return () => {
            if (unlisten) unlisten();
        };
    }, [navigate]);
    return null;
}


function App() {
    return (
        <HashRouter>
            <div className="min-h-screen w-full bg-gray-900 text-white flex flex-col">
                <NavigationHandler />
                <main className="flex-1 overflow-y-auto">
                    <Routes>
                        <Route path="/" element={<Dashboard />} />
                        <Route path="/scan" element={<Scan />} />
                        <Route path="/conn" element={<Connection />} />
                    </Routes>
                </main>

            </div>
        </HashRouter>
    )
}

export default App
