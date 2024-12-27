import os
import re

def remove_comments_and_prepend_includes(directory):
    # Define the include statements to add to the top of each file
    include_statements = """#include <stdio.h>\n#include "libft.h"\n\n"""

    try:
        for filename in os.listdir(directory):
            # Process only `.c` files
            if filename.endswith(".c"):
                file_path = os.path.join(directory, filename)
                
                # Read the file contents
                with open(file_path, "r") as file:
                    content = file.read()
                
                # Remove all comments (single-line and multi-line) using regex
                content_no_comments = re.sub(r"/\*.*?\*/|//.*?$", "", content, flags=re.DOTALL | re.MULTILINE)
                
                # Prepend the include statements to the file
                updated_content = include_statements + content_no_comments.strip() + "\n"
                
                # Write the updated content back to the file
                with open(file_path, "w") as file:
                    file.write(updated_content)
                
                print(f"Processed: {filename}")
    except Exception as e:
        print(f"Error: {e}")

# Specify the directory
directory_path = "./libft_c_tests"

# Call the function
remove_comments_and_prepend_includes(directory_path)
