import { render, screen } from '@testing-library/react';
import App from '../App';

test('renders the correct content', () => {
  render(<App />);
  expect(screen.getByText(/Hello Vite/i)).toBeInTheDocument();
});