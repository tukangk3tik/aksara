import React, { ButtonHTMLAttributes } from 'react';

type ButtonVariant = 'primary' | 'secondary' | 'outline' | 'danger' | 'success';
type ButtonSize = 'sm' | 'md' | 'lg';

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  size?: ButtonSize;
  isLoading?: boolean;
  fullWidth?: boolean;
  icon?: React.ReactNode;
  iconPosition?: 'left' | 'right';
}

const Button: React.FC<ButtonProps> = ({
  children,
  variant = 'primary',
  size = 'md',
  isLoading = false,
  fullWidth = false,
  icon,
  iconPosition = 'left',
  className = '',
  disabled,
  ...props
}) => {
  // Base classes for all buttons
  const baseClasses = 'inline-flex items-center justify-center font-medium rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2 transition-colors';
  
  // Size classes
  const sizeClasses = {
    sm: 'px-3 py-1.5 text-xs',
    md: 'px-4 py-2 text-sm',
    lg: 'px-6 py-3 text-base'
  };
  
  // Variant classes
  const variantClasses = {
    primary: 'text-white bg-amber-600 hover:bg-amber-700 focus:ring-amber-500 border border-transparent',
    secondary: 'text-white bg-gray-600 hover:bg-gray-700 focus:ring-gray-500 border border-transparent',
    outline: 'text-amber-600 bg-transparent hover:bg-amber-50 focus:ring-amber-500 border border-amber-600',
    danger: 'text-white bg-red-600 hover:bg-red-700 focus:ring-red-500 border border-transparent',
    success: 'text-white bg-green-600 hover:bg-green-700 focus:ring-green-500 border border-transparent'
  };
  
  // Width classes
  const widthClasses = fullWidth ? 'w-full' : '';
  
  // Disabled and loading states
  const stateClasses = (disabled || isLoading) 
    ? 'opacity-70 cursor-not-allowed' 
    : 'cursor-pointer';

  // Combine all classes
  const buttonClasses = `
    ${baseClasses} 
    ${sizeClasses[size]} 
    ${variantClasses[variant]} 
    ${widthClasses} 
    ${stateClasses} 
    ${className}
  `.trim();

  // Loading spinner
  const LoadingSpinner = () => (
    <svg className="animate-spin -ml-1 mr-2 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
      <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
      <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
    </svg>
  );

  return (
    <button
      className={buttonClasses}
      disabled={disabled || isLoading}
      {...props}
    >
      {isLoading && <LoadingSpinner />}
      {!isLoading && icon && iconPosition === 'left' && <span className="mr-2">{icon}</span>}
      {children}
      {!isLoading && icon && iconPosition === 'right' && <span className="ml-2">{icon}</span>}
    </button>
  );
};

export default Button;
