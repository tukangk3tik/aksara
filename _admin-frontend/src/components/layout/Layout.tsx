import React from 'react';
import { Outlet } from 'react-router-dom';
import Sidebar from './Sidebar';
import Header from './Header';
import Breadcrumb from './Breadcrumb';
import { useSidebar } from '../../contexts/SiderbarContextProps';

const Layout: React.FC = () => {
  const { isOpen } = useSidebar();

  return (
    <div className="flex h-screen bg-gray-50">
      <Sidebar />
      
      <div className={`flex-1 flex flex-col transition-all duration-300 ${
        isOpen ? 'lg:ml-64' : 'lg:ml-16'
      }`}>
        <Header />
        
        <main className="flex-1 overflow-y-auto p-4 md:p-6">
          <div className="mb-6">
            <Breadcrumb />
          </div>
          
          <div className="container mx-auto">
            <Outlet />
          </div>
        </main>
        
        <footer className="bg-white border-t border-gray-200 py-4 px-6 text-center text-sm text-gray-500">
          <p> {new Date().getFullYear()} Aksara Admin Dashboard. All rights reserved.</p>
        </footer>
      </div>
    </div>
  );
};

export default Layout;