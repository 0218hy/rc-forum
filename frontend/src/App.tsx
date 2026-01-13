import './App.css'
import { ThemeProvider } from '@mui/material'
import { theme } from './Theme'
import Navbar from "./components/Navbar"
import Feed from "./components/Feed"
import Sidebar from "./components/Sidebar"


export default function App() {

  return (
    <ThemeProvider theme={theme}>
      <Navbar/>
      <Feed />
      <Sidebar />
      <p> Hello </p>
    </ThemeProvider>
  )
}

