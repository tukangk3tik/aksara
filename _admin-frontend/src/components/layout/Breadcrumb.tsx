import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { navigationItems } from '../../routes/navConfig';
import { BreadcrumbItem } from '../../types/navigation';
import { HomeIcon, ChevronRightIcon } from '@heroicons/react/solid';

const Breadcrumb: React.FC = () => {
  const location = useLocation();
  const [breadcrumbs, setBreadcrumbs] = React.useState<BreadcrumbItem[]>([]);

  React.useEffect(() => {
    const generateBreadcrumbs = () => {
      const pathSegments = location.pathname.split('/').filter(Boolean);
      
      if (pathSegments.length === 0) {
        setBreadcrumbs([{ title: 'Aksara', path: '/' }]);
        return;
      }
      
      const breadcrumbItems: BreadcrumbItem[] = [];
      
      let currentPath = '';
      
      pathSegments.forEach((segment) => {
        currentPath += `/${segment}`;
        
        // Find matching navigation item
        let matchingItem: BreadcrumbItem | undefined;
        
        // Check top-level items
        navigationItems.forEach(item => {
          if (item.path === currentPath) {
            matchingItem = { title: item.title, path: item.path };
          }
          
          // Check children
          if (item.children) {
            item.children.forEach(child => {
              if (child.path === currentPath) {
                // Add parent if not already in breadcrumbs
                if (!breadcrumbItems.some(b => b.title === item.title)) {
                  breadcrumbItems.push({ title: item.title, path: item.path });
                }
                matchingItem = { title: child.title, path: child.path };
              }
            });
          }
        });
        
        if (matchingItem) {
          breadcrumbItems.push(matchingItem);
        }
      });
      
      setBreadcrumbs(breadcrumbItems);
    };
    
    generateBreadcrumbs();
  }, [location]);

  if (breadcrumbs.length <= 1) {
    return null;
  }

  return (
    <nav className="flex" aria-label="Breadcrumb">
      <ol className="inline-flex items-center space-x-1 md:space-x-3">
        {breadcrumbs.map((item, index) => {
          const isLast = index === breadcrumbs.length - 1;
          
          return (
            <li key={index} className="inline-flex items-center">
              {index === 0 ? (
                <Link 
                  to={item.path || '/'}
                  className="inline-flex items-center text-sm text-gray-500 hover:text-primary"
                >
                  <HomeIcon className="w-4 h-4 mr-2" />
                  {item.title}
                </Link>
              ) : (
                <div className="flex items-center">
                  <ChevronRightIcon className="w-5 h-5 text-gray-400" />
                  {isLast ? (
                    <span className="ml-1 text-sm font-medium text-gray-700 md:ml-2">
                      {item.title}
                    </span>
                  ) : (
                    <Link
                      to={item.path || '#'}
                      className="ml-1 text-sm text-gray-500 hover:text-primary md:ml-2"
                    >
                      {item.title}
                    </Link>
                  )}
                </div>
              )}
            </li>
          );
        })}
      </ol>
    </nav>
  );
};

export default Breadcrumb;