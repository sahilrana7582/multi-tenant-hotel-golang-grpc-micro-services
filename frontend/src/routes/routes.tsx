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
    ],
  },
];

export default appRoutes;
