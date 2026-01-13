import { Campaign, Groups, LocalGroceryStore, ReportProblem} from "@mui/icons-material"
import { List, Box, ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material"

const Sidebar = () => {
  return (
    <Box flex={0.5} p={1} sx={{ display: { xs: "none", sm: "block" }}} >
        <Box position = "fixed">
            <List>
                <ListItemButton href = "#announcement">
                    <ListItemIcon> <Campaign/></ListItemIcon>
                    <ListItemText primary = "Announcement" />
                </ListItemButton>

                <ListItemButton href = "#report">
                    <ListItemIcon> <ReportProblem/></ListItemIcon>
                    <ListItemText primary = "Report" />
                </ListItemButton>

                <ListItemButton href = "#marketplace">
                    <ListItemIcon> <LocalGroceryStore/></ListItemIcon>
                    <ListItemText primary = "Marketplace" />
                </ListItemButton>

                <ListItemButton href = "#openjio">
                    <ListItemIcon> <Groups/></ListItemIcon>
                    <ListItemText primary = "Open Jio" />
                </ListItemButton>
            </List>
        </Box>
    </Box>
  )
}

export default Sidebar