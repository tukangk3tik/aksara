import React, { useState, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import { useSidebar } from '../../contexts/SiderbarContextProps';
import { useAuth } from '../../contexts/AuthContext';
import { MenuIcon, BellIcon, SearchIcon, LogoutIcon } from '@heroicons/react/outline';

const Header: React.FC = () => {
  const { toggleSidebar } = useSidebar();
  const { logout } = useAuth();
  const navigate = useNavigate();
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  const toggleDropdown = () => {
    setDropdownOpen(!dropdownOpen);
  };

  // Close dropdown when clicking outside
  React.useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setDropdownOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <header className="bg-white border-b border-gray-200 h-16 flex items-center justify-between px-4 md:px-6">
      <div className="flex items-center">
        <button
          onClick={toggleSidebar}
          className="p-2 rounded-md text-gray-500 hover:bg-gray-100 lg:hidden"
        >
          <MenuIcon className="h-6 w-6" />
        </button>
      </div>
      
      <div className="flex-1 max-w-md mx-4 hidden md:block">
        <div className="relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <SearchIcon className="h-5 w-5 text-gray-400" />
          </div>
          <input
            type="text"
            className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-gray-50 placeholder-gray-500 focus:outline-none focus:ring-primary focus:border-primary sm:text-sm"
            placeholder="Search..."
          />
        </div>
      </div>
      
      <div className="flex items-center space-x-4">
        <button className="p-1 rounded-full text-gray-500 hover:bg-gray-100 relative">
          <BellIcon className="h-6 w-6" />
          <span className="absolute top-0 right-0 block h-2 w-2 rounded-full bg-red-500 ring-2 ring-white"></span>
        </button>
        
        <div className="border-l border-gray-200 h-6 mx-2"></div>
        
        <div className="relative" ref={dropdownRef}>
          <button 
            onClick={toggleDropdown}
            className="flex items-center focus:outline-none"
          >
            <img
              src="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png"
              alt="User"
              className="h-8 w-8 rounded-full"
            />
          </button>
          
          {dropdownOpen && (
            <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-10 border border-gray-200">
              <button
                onClick={handleLogout}
                className="flex items-center w-full px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              >
                <LogoutIcon className="h-4 w-4 mr-2" />
                Logout
              </button>
            </div>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;