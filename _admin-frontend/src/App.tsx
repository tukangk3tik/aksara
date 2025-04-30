import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import { SidebarProvider } from './contexts/SiderbarContextProps'
import Layout from './components/layout/Layout'
import Dashboard from './pages/Dashboard'
import School from './pages/master/School'
// import Student from './pages/master/Student'
import './App.css'

// Create placeholder components for pages that don't exist yet
const Teacher = () => <div className="p-4 bg-white rounded-lg shadow">Teacher Management Page</div>
const Users = () => <div className="p-4 bg-white rounded-lg shadow">Users Management Page</div>
const Settings = () => <div className="p-4 bg-white rounded-lg shadow">Settings Page</div>

function App() {
  return (
    <Router>
      <SidebarProvider>
        <Routes>
          <Route path="/" element={<Layout />}>
            {/* Dashboard */}
            <Route index element={<Dashboard />} />
            
            {/* Master Data Routes */}
            <Route path="master">
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
        </Routes>
      </SidebarProvider>
    </Router>
  )
}

export default App
