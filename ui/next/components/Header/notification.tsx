import Badge from '@material-ui/core/Badge'
import Menu from '@material-ui/core/Menu'
import Fade from '@material-ui/core/Fade'
import IconButton from '@material-ui/core/IconButton'
import React from 'react'
import NotificationsIcon from '@material-ui/icons/Notifications'
import { makeStyles } from '@material-ui/core/styles'

const useStyles = makeStyles(() => ({
  list: {
    margin: 0,
    padding: 0,
  },
}))

export default function Notification() {
  const classes = useStyles()

  const [anchorEl, setAnchorEl] = React.useState(null)
  const handleClick = (event?: any) => {
    setAnchorEl(event.currentTarget)
  }

  const handleClose = () => {
    setAnchorEl(null)
  }

  return (
    <IconButton
      color="inherit"
      aria-controls="simple-menu"
      aria-haspopup="true"
      onClick={handleClick}
    >
      <Badge badgeContent={4} color="secondary">
        <NotificationsIcon />
      </Badge>

      <Menu
        id="simple-menu"
        anchorEl={anchorEl}
        keepMounted
        classes={{ list: classes.list }}
        open={Boolean(anchorEl)}
        onClose={handleClose}
        TransitionComponent={Fade}
      >
        {/* Dropdown menu */}
        <div className="right-0 z-20 mt-2 overflow-hidden bg-white rounded-md shadow-lg w-80 dark:bg-gray-800">
          <div className="py-2">
            <a className="flex items-center px-4 py-3 -mx-2 transition-colors duration-200 transform border-b hover:bg-gray-100 dark:hover:bg-gray-700 dark:border-gray-700">
              <img
                className="object-cover w-8 h-8 mx-1 rounded-full"
                src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=334&q=80"
                alt="avatar"
              />
              <p className="mx-2 text-sm text-gray-600 dark:text-white">
                <span className="font-bold">Sara Salah</span> replied on the{' '}
                <span className="font-bold text-blue-500">Upload Image</span>{' '}
                artical . 2m
              </p>
            </a>
            <a className="flex items-center px-4 py-3 -mx-2 transition-colors duration-200 transform border-b hover:bg-gray-100 dark:hover:bg-gray-700 dark:border-gray-700">
              <img
                className="object-cover w-8 h-8 mx-1 rounded-full"
                src="https://images.unsplash.com/photo-1531427186611-ecfd6d936c79?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=634&q=80"
                alt="avatar"
              />
              <p className="mx-2 text-sm text-gray-600 dark:text-white">
                <span className="font-bold">Slick Net</span> start following you
                . 45m
              </p>
            </a>
            <a className="flex items-center px-4 py-3 -mx-2 transition-colors duration-200 transform border-b hover:bg-gray-100 dark:hover:bg-gray-700 dark:border-gray-700">
              <img
                className="object-cover w-8 h-8 mx-1 rounded-full"
                src="https://images.unsplash.com/photo-1450297350677-623de575f31c?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=334&q=80"
                alt="avatar"
              />
              <p className="mx-2 text-sm text-gray-600 dark:text-white">
                <span className="font-bold">Jane Doe</span> Like Your reply on{' '}
                <span className="font-bold text-blue-500">Test with TDD</span>{' '}
                artical . 1h
              </p>
            </a>
            <a className="flex items-center px-4 py-3 -mx-2 transition-colors duration-200 transform hover:bg-gray-100 dark:hover:bg-gray-700">
              <img
                className="object-cover w-8 h-8 mx-1 rounded-full"
                src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=398&q=80"
                alt="avatar"
              />
              <p className="mx-2 text-sm text-gray-600 dark:text-white">
                <span className="font-bold">Abigail Bennett</span> start
                following you . 3h
              </p>
            </a>
          </div>
          <a className="block py-2 font-bold text-center text-white bg-gray-800 dark:bg-gray-700 hover:underline">
            See all notifications
          </a>
        </div>
      </Menu>
    </IconButton>
  )
}
