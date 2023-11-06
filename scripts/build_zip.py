"""Build a zip file from a directory."""
import os
import zipfile


def build_zip(zip_filename: str, directory: str) -> None:
    """Build a zip file from a directory.

    :type zip_filename: str
    :param zip_filename: The name of the zip file to create.
    :type directory: str
    :param directory: The directory to zip.

    :rtype: None
    :return: None
    """
    with zipfile.ZipFile(zip_filename, 'w') as zf:
        for dirname, _, files in os.walk(directory):
            for filename in files:
                zf.write(os.path.join(dirname, filename), arcname=filename)


if __name__ == '__main__':
    base_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    wanted_dir = os.path.join(base_dir, "temp")
    windows_path = os.path.join(wanted_dir, "windows-amd64")
    build_zip("windows-amd64.zip", windows_path)
