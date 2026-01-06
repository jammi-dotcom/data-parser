import logging
import os
from typing import Dict, List

class Utils:
    def __init__(self, logger: logging.Logger = None):
        if logger is None:
            self.logger = logging.getLogger(__name__)
        else:
            self.logger = logger

    def read_file(self, file_path: str) -> str:
        try:
            with open(file_path, 'r') as file:
                return file.read()
        except FileNotFoundError:
            self.logger.error(f"File not found: {file_path}")
            return None

    def write_file(self, file_path: str, content: str) -> bool:
        try:
            with open(file_path, 'w') as file:
                file.write(content)
            return True
        except Exception as e:
            self.logger.error(f"Error writing file: {file_path} - {str(e)}")
            return False

    def get_env_variables(self) -> Dict[str, str]:
        return dict(os.environ)

    def parse_list(self, input_list: List[str]) -> List[str]:
        return [item.strip() for item in input_list]

def main():
    utils = Utils()
    file_path = 'example.txt'
    content = 'Hello, World!'
    utils.write_file(file_path, content)
    print(utils.read_file(file_path))

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    main()