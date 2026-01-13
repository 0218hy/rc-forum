import { styled, Toolbar, Box, AppBar, Typography, InputBase, Avatar, Menu, MenuItem} from "@mui/material";
import { Apartment } from "@mui/icons-material";
import React, { useState } from "react";

const StyledToolBar = styled(Toolbar) ({
    display: "flex",
    justifyContent: "space-between",
});

const Search = styled("div")(({theme}) => ({
    backgroundColor: "white",
    padding: "0 10px",
    borderRadius: theme.shape.borderRadius,
    width: "40%",
}));

const UserBox = styled(Box) ({
    display: "flex",
    alignItems: "center",
    gap: "10px",
});

const Navbar = () => {
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement> (null)
    const open = Boolean (anchorEl);

    const handleClick = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget)
    };
    const handleClose = () => {
        setAnchorEl(null);
    };

    return (
        <AppBar position = "sticky">
            <StyledToolBar>
                <Typography variant = "h6" sx = {{display : { xs: "none", sm: "block"}}}>
                    Hayoung
                </Typography>
                <Apartment sx={{display: { xs: "block", sm: "none"}}} />
                
                <Search>
                    <InputBase placeholder = "search for posts" />
                </Search>

                <UserBox onClick = {handleClick}>
                    <Avatar {...stringAvatar("Lee Hayoung")} />
                    <Typography variant="overline"> Hayoung</Typography>
                </UserBox>
                <Menu
                    anchorEl={anchorEl}
                    open={open}
                    onClose={handleClose}
                    anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
                    transformOrigin={{ vertical: "top", horizontal: "right" }}
                    slotProps={{
                        paper: {
                            sx: {
                                mt: 1.5,
                                minWidth: 120,
                                borderRadius: 2,
                                boxShadow: "0px 8px 16px rgba(0,0,0,0.12)",
                            },
                        },
                    }}
                >
                    <MenuItem>Profile</MenuItem>
                    <MenuItem>My posts</MenuItem>
                    <MenuItem>Logout</MenuItem>
                </Menu>
            </StyledToolBar>
        </AppBar>
    )
}

// Utility function for Avatar initial & backgroundColor (from MUI)
function stringToColor(string: string) {
    let hash = 0;
    let i;
  
    /* eslint-disable no-bitwise */
    for (i = 0; i < string.length; i += 1) {
      hash = string.charCodeAt(i) + ((hash << 5) - hash);
    }
  
    let color = '#';
  
    for (i = 0; i < 3; i += 1) {
      const value = (hash >> (i * 8)) & 0xff;
      color += `00${value.toString(16)}`.slice(-2);
    }
    /* eslint-enable no-bitwise */
  
    return color;
}

function stringAvatar(name: string) {
    return {
      sx: {
        bgcolor: stringToColor(name),
      },
      children: `${name.split(' ')[0][0]}${name.split(' ')[1][0]}`,
    };
}

export default Navbar