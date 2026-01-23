import './App.css'
import { Box, Stack, ThemeProvider, Toolbar } from '@mui/material'
import { theme } from './Theme'
import Navbar from "./components/Navbar"
import Feed from "./components/Feed"
import Sidebar from "./components/Sidebar"
import { useState } from 'react'


export default function App() {
  const [mobileOpen, setMobileOpen] = useState(false);
  const handleMenuClick = () => setMobileOpen(true);
  const handleSideBarClose = () => setMobileOpen(false);

  return (
    <ThemeProvider theme={theme}>
      <Box bgcolor= {"background.default"} color = {"text.primary"}>
        <Navbar onMenuClick={handleMenuClick}/>

        <Toolbar />
        
        <Box sx={{ display: "flex" }}>
          <Sidebar mobileOpen={mobileOpen} onClose={handleSideBarClose} />
          <Box
            sx={{
              flexGrow: 1,
              p: 2,
              ml: { sm: "240px" },
            }}
          >
            <Feed />
          </Box>
        </Box>
      </Box>
    </ThemeProvider>
  )
}

