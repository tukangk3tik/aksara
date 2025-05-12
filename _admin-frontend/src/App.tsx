import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import { SidebarProvider } from './contexts/SiderbarContextProps'
import { AuthProvider, useAuth } from './contexts/AuthContext'
import Layout from './components/layout/Layout'
import Dashboard from './pages/Dashboard'
import School from './pages/master/School'
import Office from './pages/master/Office'
import Login from './pages/Login'
import ProtectedRoute from './components/ProtectedRoute'
// import Student from './pages/master/Student'
import './App.css'

// Create placeholder components for pages that don't exist yet
const Teacher = () => <div className="p-4 bg-white rounded-lg shadow">Teacher Management Page</div>
const Users = () => <div className="p-4 bg-white rounded-lg shadow">Users Management Page</div>
const Settings = () => <div className="p-4 bg-white rounded-lg shadow">Settings Page</div>

// AppRoutes component to handle routing after authentication check
const AppRoutes = () => {
  return (
    <Routes>
      {/* Public routes */}
      <Route path="/login" element={<Login />} />
      
      {/* Protected routes */}
      <Route element={<ProtectedRoute />}>
        <Route path="/" element={<Layout />}>
          {/* Dashboard */}
          <Route index element={<Dashboard />} />
          
          {/* Master Data Routes */}
          <Route path="master">
            <Route path="office" element={<Office />} />
            <Route path="school" element={<School />} />
            {/* <Route path="student" element={<Student />} /> */}
            <Route path="teacher" element={<Teacher />} />
            <Route index element={<Navigate to="/master/school" replace />} />
          </Route>
          
          {/* Users */}
          <Route path="users" element={<Users />} />
          
          {/* Settings */}
          <Route path="settings" element={<Settings />} />
          
          {/* Catch all - redirect to dashboard */}
          <Route path="*" element={<Navigate to="/" replace />} />
        </Route>
      </Route>
    </Routes>
  );
};

function App() {
  return (
    <Router>
      <AuthProvider>
        <SidebarProvider>
          <AppRoutes />
        </SidebarProvider>
      </AuthProvider>
    </Router>
  )
}

export default App
