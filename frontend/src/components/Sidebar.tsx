import { Campaign, FiberNew, Groups, LocalGroceryStore, ReportProblem} from "@mui/icons-material"
import { List, Box, ListItemButton, ListItemIcon, ListItemText, Toolbar, Drawer } from "@mui/material"

const drawerWidth = 240;

type Props = {
  mobileOpen: boolean;
  onClose: () => void;
};

export default function Sidebar({ mobileOpen, onClose }: Props) {
  const drawerContent = (
    <Box>
      <List>
        <ListItemButton href="#announcement" onClick={onClose}>
          <ListItemIcon><Campaign /></ListItemIcon>
          <ListItemText primary="Announcement" />
        </ListItemButton>

        <ListItemButton href="#report" onClick={onClose}>
          <ListItemIcon><ReportProblem /></ListItemIcon>
          <ListItemText primary="Report" />
        </ListItemButton>

        <ListItemButton href="#marketplace" onClick={onClose}>
          <ListItemIcon><LocalGroceryStore /></ListItemIcon>
          <ListItemText primary="Marketplace" />
        </ListItemButton>

        <ListItemButton href="#openjio" onClick={onClose}>
          <ListItemIcon><Groups /></ListItemIcon>
          <ListItemText primary="Open Jio" />
        </ListItemButton>

        <ListItemButton href="#create" onClick={onClose}>
          <ListItemIcon><FiberNew /></ListItemIcon>
          <ListItemText primary="Create New Post!" />
        </ListItemButton>
      </List>
    </Box>
  );

  return (
    <>
    {/* mobile */}
      <Drawer
        variant="temporary"
        open={mobileOpen}
        onClose={onClose}
        ModalProps={{ keepMounted: true }}
        sx={{
          display: { xs: "block", sm: "none" },
          "& .MuiDrawer-paper": {
            width: drawerWidth,
          },
        }}
      >
        {drawerContent}
      </Drawer>

     {/* bigger than mobile */}
      <Drawer
        variant="permanent"
        sx={{
          display: { xs: "none", sm: "block" },
          "& .MuiDrawer-paper": {
            width: drawerWidth,
            boxSizing: "border-box",
            top: '64px', // Navbar height
            height: 'calc(100% - 64px)',
          },
        }}
        open
      >
        {drawerContent}
      </Drawer>
    </>
  );
}
