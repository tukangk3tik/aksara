import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useSidebar } from '../../contexts/SiderbarContextProps';
import { navigationItems } from '../../routes/navConfig';
import { 
  HomeIcon, 
  UserGroupIcon, 
  CogIcon, 
  DatabaseIcon,
  ChevronDownIcon,
  ChevronRightIcon,
  MenuIcon,
  XIcon
} from '@heroicons/react/outline';
import logoImage from '../../assets/images/logo.png';

const iconMap: Record<string, React.ReactElement> = {
  HomeIcon: <HomeIcon className="h-5 w-5 text-gray-500" />,
  UserGroupIcon: <UserGroupIcon className="h-5 w-5 text-gray-500" />,
  CogIcon: <CogIcon className="h-5 w-5 text-gray-500" />,
  DatabaseIcon: <DatabaseIcon className="h-5 w-5 text-gray-500" />,
};

const Sidebar: React.FC = () => {
  const { isOpen, toggleSidebar, expandedMenus, toggleMenu } = useSidebar();
  const location = useLocation();

  const renderNavItems = (items: typeof navigationItems) => {
    return items.map((item) => {
      const isActive = location.pathname === item.path;
      const hasChildren = item.children && item.children.length > 0;
      const isExpanded = expandedMenus.includes(item.id);
      
      // Check if any child is active
      const isChildActive = hasChildren && item.children?.some(
        child => location.pathname === child.path
      );

      return (
        <div key={item.id} className="mb-1">
          {hasChildren ? (
            <>
              <button
                onClick={() => toggleMenu(item.id)}
                className={`w-full flex items-center px-2 py-2 text-sm rounded-md transition-colors ${
                  isChildActive ? 'bg-primary-light/20 text-primary-dark' : 'hover:bg-orange-100'
                }`}
              >
                <div className="flex items-center">
                  {item.icon && iconMap[item.icon]}
                  <span className={`ml-3 ${isOpen ? 'block' : 'hidden'}`}>{item.title}</span>
                </div>
                <div className={`${isOpen ? 'block' : 'hidden'} pl-4`}>
                  {isExpanded ? 
                    <ChevronDownIcon className="h-4 w-4" /> : 
                    <ChevronRightIcon className="h-4 w-4" />
                  }
                </div>
              </button>
              
              {isExpanded && isOpen && (
                <div className="pl-10 mt-1 space-y-1">
                  {item.children?.map((child) => {
                    const isChildItemActive = location.pathname === child.path;
                    
                    return (
                      <Link
                        key={child.id}
                        to={child.path || '#'}
                        className={`flex items-center px-4 py-2 text-sm rounded-md ${
                          isChildItemActive 
                            ? 'bg-orange-200 text-primary-dark font-medium' 
                            : 'hover:bg-orange-100'
                        }`}
                      >
                        {child.title}
                      </Link>
                    );
                  })}
                </div>
              )}
            </>
          ) : (
            <Link
              to={item.path || '#'}
              className={`w-full flex items-center px-2 py-2 text-sm rounded-md transition-colors ${
                isActive ? 'bg-orange-200 text-primary-dark font-medium' : 'hover:bg-orange-100'
              }`}
            >
              <div className="flex items-center justify-center w-5 h-5">
                {item.icon && iconMap[item.icon]}
              </div>
              <span className={`ml-3 ${isOpen ? 'block' : 'hidden'}`}>{item.title}</span>
            </Link>
          )}
        </div>
      );
    });
  };

  return (
    <div 
      className={`fixed inset-y-0 left-0 z-30 bg-white border-r border-gray-200 transition-all duration-300 ${
        isOpen ? 'w-64' : 'w-16'
      }`}
    >
      <div className="flex items-center justify-between h-16 px-4 border-b border-gray-200">
        <div className="flex items-center">
          {isOpen ? <img 
            src={logoImage} 
            alt="Aksara Logo" 
            className="h-8 w-8"
          /> : ''}
          {isOpen && <span className="ml-2 text-lg font-semibold">Aksara Admin</span>}
        </div>
        <button 
          onClick={toggleSidebar}
          className="p-1 rounded-md hover:bg-orange-100"
        >
          {isOpen ? 
            <XIcon className="h-5 w-5" /> : 
            <MenuIcon className="h-5 w-5" />
          }
        </button>
      </div>
      
      <div className="p-4">
        <div className="flex items-center mb-6">
          <div className="relative">
            <img 
              src="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png" 
              alt="User" 
              className="h-10 w-10 rounded-full"
            />
            <span className="absolute bottom-0 right-0 block h-2.5 w-2.5 rounded-full bg-green-400 border-2 border-white"></span>
          </div>
          {isOpen && (
            <div className="ml-3">
              <p className="text-sm font-medium">Yanetta</p>
              <p className="text-xs text-gray-500">yanetta@gmail.com</p>
            </div>
          )}
        </div>
        
        <nav className="space-y-1">
          {renderNavItems(navigationItems)}
        </nav>
      </div>
    </div>
  );
};

export default Sidebar;