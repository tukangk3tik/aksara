
export interface SelectOption {
  id: number;
  name: string;
}

export interface SelectOptionLocationData {
  label: string;
  values: SelectOption[];
  isLoading: boolean;
  selectedValue: SelectOption | null;
  searchTerm: string;
  filteredOptions: SelectOption[];
  fieldError: string | null;
  placeholder?: string;
}

export interface SelectOptionProps {
  data: SelectOptionLocationData;
  onChange: (value: SelectOption) => void;
  handleQueryChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleClearSearch: () => void;
}