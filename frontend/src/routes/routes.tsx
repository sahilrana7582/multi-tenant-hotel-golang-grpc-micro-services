import Layout from "../components/Layout/Layout";
import Login from "../pages/login/Login";


const appRoutes = [
  // 🔓 Public Route
  {
    path: '/login',
    element: <Login />,
  },

  // 🔒 Protected Routes (Layout with sidebar/header)
  {
    path: '/',
    element: <Layout />,
    children: [
      { path: '', element: <h1>LALA DASH</h1> },
      { path: 'rooms', element: <h1>Rooms</h1> },
      { path: 'housekeeping', element: <h1>HK</h1> },
      { path: 'messages', element: <h1>messages</h1> },
      { path: 'kitchen', element: <h1>kitchen</h1> },
      { path: 'frontdesk', element: <h1>frontdesk</h1> },
    ],
  },
];

export default appRoutes;
